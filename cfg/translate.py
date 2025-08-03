#!/usr/bin/env python

import argparse
import os
import tempfile
import subprocess
from io import StringIO
from pathlib import Path

# The Ruamel Yaml loader supports "round trip" loading/saving, where comments/etc can be maintained
from ruamel.yaml import YAML

yaml = YAML()
# Preserve quotes avoid unecessary quote style replacements
yaml.preserve_quotes = True

FOLDER = os.path.dirname(__file__)
BAD_KEYS = ["_name_", "_attrs_order_", "_required_"]


def generate_focused_diff(original_content, translated_content, changes, output_file=None):
    """Generate a git-style diff only for changed components"""
    with tempfile.TemporaryDirectory() as tmpdir:
        # Create temp files
        orig_file = Path(tmpdir) / "original.yaml"
        new_file = Path(tmpdir) / "translated.yaml"

        # Load both specs
        yaml = YAML()
        yaml.preserve_quotes = True  # Preserve the original quoting style
        orig_spec = yaml.load(original_content)
        new_spec = yaml.load(translated_content)

        # Create minimal specs while preserving original structure
        focused_orig = {
            "openapi": orig_spec.get("openapi", "3.0.0"),
            "info": orig_spec.get("info", {}),
            "paths": {},
            "components": {
                "schemas": {},
                "responses": {},
                "parameters": {},
                "requestBodies": {},
            }
        }

        focused_new = {
            "openapi": new_spec.get("openapi", "3.0.0"),
            "info": new_spec.get("info", {}),
            "paths": {},
            "components": {
                "schemas": {},
                "responses": {},
                "parameters": {},
                "requestBodies": {},
            }
        }

        # Copy paths from original content
        for path in changes['paths']:
            if path in orig_spec.get('paths', {}):
                focused_orig['paths'][path] = orig_spec['paths'][path]
            if path in new_spec.get('paths', {}):
                focused_new['paths'][path] = new_spec['paths'][path]

        # Copy components that are referenced
        for comp_type, items in changes['components'].items():
            for name in items:
                if comp_type in orig_spec.get('components', {}) and name in orig_spec['components'][comp_type]:
                    focused_orig['components'][comp_type][name] = orig_spec['components'][comp_type][name]
                if comp_type in new_spec.get('components', {}) and name in new_spec['components'][comp_type]:
                    focused_new['components'][comp_type][name] = new_spec['components'][comp_type][name]

        # Write focused specs to temp files
        with open(orig_file, 'w') as f:
            yaml.dump(focused_orig, f)

        with open(new_file, 'w') as f:
            yaml.dump(focused_new, f)

        # Generate diff using git
        try:
            subprocess.run(['git', 'init'], cwd=tmpdir, capture_output=True, check=True)
            subprocess.run(['git', 'config', 'user.email', 'temp@example.com'], cwd=tmpdir, capture_output=True,
                           check=True)
            subprocess.run(['git', 'config', 'user.name', 'Temporary User'], cwd=tmpdir, capture_output=True,
                           check=True)

            # Add original file
            subprocess.run(['git', 'add', str(orig_file)], cwd=tmpdir, capture_output=True, check=True)
            subprocess.run(['git', 'commit', '-m', 'Original'], cwd=tmpdir, capture_output=True, check=True)

            # Replace with new content and show diff
            new_file.rename(orig_file)

            # Use git diff with more context
            result = subprocess.run(
                ['git', 'diff', '--no-color', '-U10'],  # Show 10 lines of context
                cwd=tmpdir,
                capture_output=True,
                text=True,
                check=True
            )

            if output_file:
                Path(output_file).write_text(result.stdout)
                print(f"Diff saved to: {output_file}")
            else:
                if result.stdout:
                    print(result.stdout)
                else:
                    print("No differences found between original and translated specs")

        except subprocess.CalledProcessError as e:
            print(f"Error generating diff: {e}")
            print(f"stdout: {e.stdout}")
            print(f"stderr: {e.stderr}")


def ensure_valid_schema(node):
    """Ensures the schema object has valid OpenAPI structure"""
    if isinstance(node, dict):
        # Handle type field validation
        if "type" in node:
            # Handle array of types (e.g., ['string', 'null'])
            if isinstance(node["type"], list):
                # Check for nullable type
                if "null" in node["type"]:
                    node["nullable"] = True
                    node["type"] = next(t for t in node["type"] if t != "null")
                else:
                    node["type"] = node["type"][0]
            # Handle object type
            elif isinstance(node["type"], dict):
                node["type"] = "object"
            # Handle other non-string types
            elif not isinstance(node["type"], str):
                if isinstance(node["type"], (int, float)):
                    node["type"] = "number"
                elif isinstance(node["type"], bool):
                    node["type"] = "boolean"
                else:
                    node["type"] = "string"

        # Handle default values based on type
        if "type" in node and "default" in node:
            if node["type"] == "object" and not isinstance(node["default"], dict):
                node["default"] = {}
            elif node["type"] == "array" and not isinstance(node["default"], list):
                node["default"] = []
            elif node["type"] == "boolean" and not isinstance(node["default"], bool):
                node["default"] = False
            elif node["type"] == "integer" and not isinstance(node["default"], int):
                node["default"] = 0
            elif node["type"] == "number" and not isinstance(node["default"], (int, float)):
                node["default"] = 0.0
            elif node["type"] == "string" and not isinstance(node["default"], str):
                node["default"] = ""

        # Remove invalid properties for non-array types
        if "type" not in node or node.get("type") != "array":
            if "items" in node and not node.get("anyOf"):
                del node["items"]

        # Handle array type schemas
        if node.get("type") == "array":
            if "items" not in node:
                node["items"] = {"type": "string"}
            elif not isinstance(node["items"], dict):
                node["items"] = ensure_valid_schema_type(node["items"])
            else:
                ensure_valid_schema(node["items"])

        # Handle anyOf, allOf, oneOf
        for key in ["anyOf", "allOf", "oneOf"]:
            if key in node and isinstance(node[key], list):
                for item in node[key]:
                    if isinstance(item, dict):
                        ensure_valid_schema(item)

        # Special handling for properties object
        if "properties" in node:
            properties = node["properties"]
            if isinstance(properties, dict):
                for prop_name, prop_value in list(properties.items()):
                    if isinstance(prop_value, dict) and any(k in prop_value for k in ["type", "anyOf", "allOf", "oneOf", "$ref", "enum"]):
                        ensure_valid_schema(prop_value)
                    else:
                        properties[prop_name] = ensure_valid_schema_type(prop_value)

        # Ensure object type is set when properties exist
        if "properties" in node and "type" not in node:
            node["type"] = "object"

        # Recursively process all dictionary values
        for k, v in node.items():
            if k not in ["properties", "items", "anyOf", "allOf", "oneOf"] and isinstance(v, dict):
                ensure_valid_schema(v)
            elif isinstance(v, list) and k not in ["anyOf", "allOf", "oneOf"]:
                for item in v:
                    if isinstance(item, dict):
                        ensure_valid_schema(item)


def ensure_valid_schema_type(value):
    """Convert a value to a valid OpenAPI schema object"""
    if isinstance(value, dict):
        if any(k in value for k in ["type", "anyOf", "allOf", "oneOf", "$ref", "enum"]):
            # It's already a schema object, just ensure it's valid
            schema = value.copy()
            ensure_valid_schema(schema)
            return schema
        else:
            # Convert to object schema
            return {"type": "object", "properties": value}
    elif isinstance(value, list):
        # For arrays, ensure we have a valid items schema
        return {
            "type": "array",
            "items": (ensure_valid_schema_type(value[0]) if value else {"type": "string"})
        }
    elif isinstance(value, bool):
        return {"type": "boolean"}
    elif isinstance(value, (int, float)):
        return {"type": "number"}
    else:
        return {"type": "string"}


def strip_bad_keys(node) -> None:
    """Strips keys that openapi doesn't expect"""
    for k in list(node.keys()):
        v = node[k]
        if k in BAD_KEYS:
            del node[k]
            continue
        elif isinstance(v, dict):
            strip_bad_keys(v)
            ensure_valid_schema(v)
        elif isinstance(v, list):
            for x in v:
                if isinstance(x, dict):
                    strip_bad_keys(x)
                    ensure_valid_schema(x)

def translate_ref_inplace(cfg, cfg_orig, ref: str, collect_changes=None) -> None:
    print(f"Translating ref: {ref}")
    assert isinstance(ref, str) and ref.startswith("#/"), f"Invalid ref: {ref}"
    ref_type, ref_subtype, ref_name = ref[2:].split("/")

    # Store original component state before translation if collecting changes
    if collect_changes is not None and ref_type == "components":
        collect_changes['components'][ref_subtype][ref_name] = cfg_orig[ref_type][ref_subtype][ref_name]

    ref_obj = cfg_orig[ref_type][ref_subtype][ref_name]
    strip_bad_keys(ref_obj)
    ensure_valid_schema(ref_obj)
    cfg[ref_type][ref_subtype][ref_name] = ref_obj


def translate_refs_inplace(cfg, cfg_orig, node, collect_changes=None) -> None:
    for k in list(node.keys()):
        v = node[k]
        if k == "$ref":
            translate_ref_inplace(cfg, cfg_orig, v, collect_changes)
        elif isinstance(v, dict):
            translate_refs_inplace(cfg, cfg_orig, v, collect_changes)


def translate_route_inplace(cfg, cfg_orig, route: str, collect_changes=None) -> None:
    """
        Translate a route and optionally collect what was changed.

        Args:
            cfg: Target config to update
            cfg_orig: Source config to read from
            route: Route to translate
            collect_changes: Optional dict to collect changes made {component_type: {name: value}}
        """
    assert route.startswith("/"), f"All routes must start with '/', got {route}"
    assert route in cfg_orig["paths"], f"Route {route} not found in (original) route paths"

    # Initialize change collection if requested
    changes = collect_changes if collect_changes is not None else {}
    changes.setdefault('paths', {})
    changes.setdefault('components', {
        'schemas': {},
        'responses': {},
        'parameters': {},
        'requestBodies': {},
    })

    # Store original route state before translation
    if collect_changes is not None:
        changes['paths'][route] = cfg_orig["paths"][route]

    translate_refs_inplace(cfg, cfg_orig, cfg_orig["paths"][route], changes)
    strip_bad_keys(cfg_orig["paths"][route])
    cfg["paths"][route] = cfg_orig["paths"][route]


def main() -> None:
    """Translate an original apidocs yaml into a curated one, with only specific routes

    For example, to update bluefin.yaml with /interface and /pool routes (from bluefin_original.yaml):
      ./cfg/translate.py bluefin /interface /pool
    """
    parser = argparse.ArgumentParser()
    parser.add_argument('release', choices=["angelfish", "bluefin", "fangtooth"], help="Release name: 'angelfish', 'bluefin' or 'fangtooth'")
    parser.add_argument('routes', nargs='*', help="Routes to replace using original schema")
    parser.add_argument('--diff', action='store_true',
                        help="Generate diff between original and translated specs")
    parser.add_argument('--diff-file', type=str,
                        help="File to save the diff to (defaults to stdout)")
    args = parser.parse_args()

    translated_file = f"cfg/{args.release}.yaml"
    original_file = f"cfg/{args.release}_original.yaml"

    if not os.path.exists(original_file):
        raise FileNotFoundError(f"Original file not found: {original_file}")

    # Store original content if diff is requested
    original_content = None
    if args.diff:
        with open(original_file, 'r') as f:
            original_content = f.read()

    with open(f"cfg/{args.release}_original.yaml", "r") as file:
        cfg_orig = yaml.load(file)

    if os.path.exists(translated_file):
        with open(translated_file, "r") as file:
            cfg = yaml.load(file)
    else:
        # Create a new base structure with minimum required OpenAPI fields
        cfg = {
            "openapi": cfg_orig.get("openapi", "3.0.0"),
            "info": cfg_orig.get("info", {"title": "", "version": ""}),
            "paths": {},
            "components": {
                "schemas": {},
                "responses": {},
                "parameters": {},
                "examples": {},
                "requestBodies": {},
                "headers": {},
                "securitySchemes": {},
                "links": {},
                "callbacks": {}
            }
        }

    # Collect changes if diff is requested
    changes = {
        'paths': {},
        'components': {
            'schemas': {},
            'responses': {},
            'parameters': {},
            'requestBodies': {},
        }
    } if args.diff else None

    for route in args.routes:
        translate_route_inplace(cfg, cfg_orig, route, changes)

    with open(translated_file, "w") as file:
        yaml.dump(cfg, file)

    if args.diff:
        with open(translated_file, "r") as file:
            translated_content = file.read()
        generate_focused_diff(original_content, translated_content, changes, args.diff_file)


if __name__ == "__main__":
    main()
