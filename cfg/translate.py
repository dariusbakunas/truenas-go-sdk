#!/usr/bin/env python

import argparse
import os
# The Ruamel Yaml loader supports "round trip" loading/saving, where comments/etc can be maintained
from ruamel.yaml import YAML

yaml = YAML()
# Preserve quotes avoid unecessary quote style replacements
yaml.preserve_quotes = True

FOLDER = os.path.dirname(__file__)
BAD_KEYS = ["_name_", "_attrs_order_", "_required_"]

def strip_bad_keys(node) -> None:
    """Strips keys that openapi doesn't expect"""
    for k in list(node.keys()):
        v = node[k]
        if k in BAD_KEYS:
            del node[k]
            continue
        elif isinstance(v, dict):
            strip_bad_keys(v)
        elif isinstance(v, list):
            for x in v:
                if isinstance(x, dict):
                    strip_bad_keys(x)

def translate_ref_inplace(cfg, cfg_orig, ref: str) -> None:
    print(f"Translating ref: {ref}")
    assert isinstance(ref, str) and ref.startswith("#/"), f"Invalid ref: {ref}"
    ref_type, ref_subtype, ref_name = ref[2:].split("/")
    ref_obj = cfg_orig[ref_type][ref_subtype][ref_name]
    strip_bad_keys(ref_obj)
    cfg[ref_type][ref_subtype][ref_name] = ref_obj


def translate_refs_inplace(cfg, cfg_orig, node) -> None:
    for k in list(node.keys()):
        v = node[k]
        if k == "$ref":
            translate_ref_inplace(cfg, cfg_orig, v)
        elif isinstance(v, dict):
            translate_refs_inplace(cfg, cfg_orig, v)


def translate_route_inplace(cfg, cfg_orig, route: str) -> None:
    assert route.startswith("/"), f"All routes must start with '/', got {route}"
    assert route in cfg_orig["paths"], f"Route {route} not found in (original) route paths"
    translate_refs_inplace(cfg, cfg_orig, cfg_orig["paths"][route])
    strip_bad_keys(cfg_orig["paths"][route])
    cfg["paths"][route] = cfg_orig["paths"][route]


def main() -> None:
    """Translate an original apidocs yaml into a curated one, with only specific routes

    For example, to update bluefin.yaml with /interface and /pool routes (from bluefin_original.yaml):
      ./cfg/translate.py bluefin /interface /pool
    """
    parser = argparse.ArgumentParser()
    parser.add_argument('release', choices=["angelfish", "bluefin"], help="Release name: 'angelfish' or 'bluefin'")
    parser.add_argument('routes', nargs='*', help="Routes to replace using original schema")
    args = parser.parse_args()

    with open(f"cfg/{args.release}.yaml", "r") as file:
        cfg = yaml.load(file)

    with open(f"cfg/{args.release}_original.yaml", "r") as file:
        cfg_orig = yaml.load(file)

    for route in args.routes:
        translate_route_inplace(cfg, cfg_orig, route)

    with open(f"cfg/{args.release}.yaml", "w") as file:
        yaml.dump(cfg, file)

if __name__ == "__main__":
    main()
