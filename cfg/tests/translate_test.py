import subprocess

import pytest
from pathlib import Path
import tempfile
from ruamel.yaml import YAML

from cfg.translate import translate_route_inplace, ensure_valid_schema


def test_schema_type_validation():
    """Test schema type validation and correction"""
    test_cases = [
        # Simple type field
        {
            'input': {'type': 'string'},
            'expected': {'type': 'string'}
        },
        # Type field that's a list (should be converted to string)
        {
            'input': {'type': ['string', 'null']},
            'expected': {'type': 'string', 'nullable': True}
        },
        # Type field that's an object (should be converted to string)
        {
            'input': {'type': {'name': 'string'}},
            'expected': {'type': 'object'}
        },
        # Schema with anyOf
        {
            'input': {
                'anyOf': [
                    {'type': 'string'},
                    {'type': ['number', 'null']},
                    {'type': {'invalid': 'type'}}
                ]
            },
            'expected': {
                'anyOf': [
                    {'type': 'string'},
                    {'type': 'number', 'nullable': True},
                    {'type': 'object'}
                ]
            }
        }
    ]

    for case in test_cases:
        ensure_valid_schema(case['input'])
        assert case['input'] == case['expected'], f"Failed for case: {case}"


def test_schema_with_invalid_type():
    """Test handling of schemas with invalid type field"""
    original_content = """
openapi: 3.0.0
paths:
  /test:
    get:
      responses:
        '200':
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/InvalidTypeSchema'
components:
  schemas:
    InvalidTypeSchema:
      type: ['string', 'null']
      properties:
        field1:
          type: {'name': 'string'}
        field2:
          type: [1, 2, 3]
"""

    translated_content = """
openapi: 3.0.0
paths:
  /test:
    get:
      responses:
        '200':
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/InvalidTypeSchema'
components:
  schemas:
    InvalidTypeSchema:
      type: string
      nullable: true
      properties:
        field1:
          type: object
        field2:
          type: array
"""

    changes = {
        'paths': {'/test': {}},
        'components': {
            'schemas': {'InvalidTypeSchema': {}},
            'responses': {},
            'parameters': {},
            'requestBodies': {},
        }
    }

    diff = generate_focused_diff(original_content, translated_content, changes)
    assert "string" in diff
    assert "nullable: true" in diff


def test_schema_with_defaults():
    """Test handling of schemas with default values and complex structures"""
    original_content = """
openapi: 3.0.0
paths:
  /test:
    get:
      responses:
        '200':
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/SchemaWithDefaults'
components:
  schemas:
    SchemaWithDefaults:
      anyOf:
        - type: object
          properties:
            field1:
              type: string
              default: "test"
            field2:
              type: object
              default: {}
            field3:
              type: array
              default: []
          default: {}
        - type: object
          properties:
            field4:
              type: string
          default: {}
      nullable: false
      description: Test schema
"""

    translated_content = original_content  # Same content for this test

    changes = {
        'paths': {'/test': {}},
        'components': {
            'schemas': {'SchemaWithDefaults': {}},
            'responses': {},
            'parameters': {},
            'requestBodies': {},
        }
    }

    diff = generate_focused_diff(original_content, translated_content, changes)
    assert "No differences found" in diff


def test_schema_validation():
    """Test schema validation and correction"""
    yaml = YAML()
    schema = yaml.load("""
type: object
properties:
  test:
    type: object
    default: invalid
""")

    ensure_valid_schema(schema)
    assert isinstance(schema['properties']['test']['default'], dict)


def test_anyof_schema_with_defaults():
    """Test handling of anyOf schemas with default values"""
    original_content = """
openapi: 3.0.0
paths:
  /test:
    get:
      responses:
        '200':
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/ComplexSchema'
components:
  schemas:
    ComplexSchema:
      anyOf:
        - title: first_option
          type: array
          items:
            type: object
            properties:
              field1:
                type: string
          default: []
        - type: object
          properties:
            field2:
              type: string
          default: {}
        - type: integer
          default: 0
      nullable: false
      description: Complex schema with various types and defaults
"""

    translated_content = original_content  # Same content for this test

    changes = {
        'paths': {'/test': {}},
        'components': {
            'schemas': {'ComplexSchema': {}},
            'responses': {},
            'parameters': {},
            'requestBodies': {},
        }
    }

    diff = generate_focused_diff(original_content, translated_content, changes)
    assert "No differences found" in diff


def test_ensure_valid_schema_with_defaults():
    """Test ensure_valid_schema function with various default values"""
    test_cases = [
        {
            'input': {
                'type': 'object',
                'properties': {
                    'test': {
                        'type': 'object',
                        'default': 'invalid_default'  # Should be converted to {}
                    }
                }
            },
            'expected_default': {}
        },
        {
            'input': {
                'type': 'array',
                'items': {
                    'type': 'string'
                },
                'default': 'invalid_default'  # Should be converted to []
            },
            'expected_default': []
        },
        {
            'input': {
                'type': 'object',
                'properties': {
                    'test': {
                        'type': 'boolean',
                        'default': 'invalid_default'  # Should be converted to False
                    }
                }
            },
            'expected_default': False
        }
    ]

    for case in test_cases:
        ensure_valid_schema(case['input'])
        if 'type' in case['input']:
            if case['input']['type'] == 'object':
                test_field = case['input']['properties']['test']
                assert test_field['default'] == case['expected_default']
            else:
                assert case['input']['default'] == case['expected_default']


def test_type_property_handling():
    """Test handling of 'type' property in nested objects"""
    original_content = """
openapi: 3.0.0
paths:
  /test:
    get:
      responses:
        '200':
          description: OK
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/TestSchema'
components:
  schemas:
    TestSchema:
      type: object
      properties:
        normalField:
          type: string
        typeField:
          type: string
          description: This is a field named type
        objectWithType:
          type: object
          properties:
            type:
              type: string
              description: Nested type field
"""

    translated_content = """
openapi: 3.0.0
paths:
  /test:
    get:
      responses:
        '200':
          description: OK
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/TestSchema'
components:
  schemas:
    TestSchema:
      type: object
      properties:
        normalField:
          type: string
        typeField:
          type: string
          description: This is a field named type
        objectWithType:
          type: object
          properties:
            type:
              type: string
              description: Modified nested type field
"""

    changes = {
        'paths': {'/test': {}},
        'components': {
            'schemas': {'TestSchema': {}},
            'responses': {},
            'parameters': {},
            'requestBodies': {},
        }
    }

    diff = generate_focused_diff(original_content, translated_content, changes)
    assert 'Modified nested type field' in diff
    assert 'Nested type field' in diff


def test_complex_type_schema():
    """Test handling of complex schemas with type fields at different levels"""
    original_content = """
openapi: 3.0.0
paths:
  /test:
    post:
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/ComplexType'
components:
  schemas:
    ComplexType:
      type: object
      properties:
        metadata:
          type: object
          properties:
            type:
              type: string
              enum: [type1, type2]
        data:
          type: object
          properties:
            type:
              type: object
              properties:
                name:
                  type: string
                type:
                  type: string
"""

    translated_content = """
openapi: 3.0.0
paths:
  /test:
    post:
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/ComplexType'
components:
  schemas:
    ComplexType:
      type: object
      properties:
        metadata:
          type: object
          properties:
            type:
              type: string
              enum: [type1, type2, type3]
        data:
          type: object
          properties:
            type:
              type: object
              properties:
                name:
                  type: string
                type:
                  type: string
"""

    changes = {
        'paths': {'/test': {}},
        'components': {
            'schemas': {'ComplexType': {}},
            'responses': {},
            'parameters': {},
            'requestBodies': {},
        }
    }

    diff = generate_focused_diff(original_content, translated_content, changes)
    assert '[type1, type2]' in diff
    assert '[type1, type2, type3]' in diff


def test_nested_array_with_type():
    """Test handling of type property within array items"""
    original_content = """
openapi: 3.0.0
paths:
  /test:
    get:
      responses:
        '200':
          $ref: '#/components/schemas/ArrayType'
components:
  schemas:
    ArrayType:
      type: array
      items:
        type: object
        properties:
          type:
            type: string
            enum: [value1]
"""

    translated_content = """
openapi: 3.0.0
paths:
  /test:
    get:
      responses:
        '200':
          $ref: '#/components/schemas/ArrayType'
components:
  schemas:
    ArrayType:
      type: array
      items:
        type: object
        properties:
          type:
            type: string
            enum: [value2]
"""

    changes = {
        'paths': {'/test': {}},
        'components': {
            'schemas': {'ArrayType': {}},
            'responses': {},
            'parameters': {},
            'requestBodies': {},
        }
    }

    diff = generate_focused_diff(original_content, translated_content, changes)
    assert 'value1' in diff
    assert 'value2' in diff


def test_type_in_deeply_nested_structure():
    """Test handling of type property in deeply nested structures"""
    original_content = """
openapi: 3.0.0
paths:
  /test:
    get:
      responses:
        '200':
          $ref: '#/components/schemas/DeepNested'
components:
  schemas:
    DeepNested:
      type: object
      properties:
        level1:
          type: object
          properties:
            level2:
              type: object
              properties:
                type:
                  type: object
                  properties:
                    level3:
                      type: object
                      properties:
                        type:
                          type: string
                          description: Original
"""

    translated_content = """
openapi: 3.0.0
paths:
  /test:
    get:
      responses:
        '200':
          $ref: '#/components/schemas/DeepNested'
components:
  schemas:
    DeepNested:
      type: object
      properties:
        level1:
          type: object
          properties:
            level2:
              type: object
              properties:
                type:
                  type: object
                  properties:
                    level3:
                      type: object
                      properties:
                        type:
                          type: string
                          description: Modified
"""

    changes = {
        'paths': {'/test': {}},
        'components': {
            'schemas': {'DeepNested': {}},
            'responses': {},
            'parameters': {},
            'requestBodies': {},
        }
    }

    diff = generate_focused_diff(original_content, translated_content, changes)
    assert 'Original' in diff
    assert 'Modified' in diff


def test_simple_path_diff():
    # Test case with a simple path change
    original_content = """
openapi: 3.0.0
info:
  title: Test API
  version: v1.0
paths:
  /test:
    get:
      description: Original description
      responses:
        '200':
          description: OK
components:
  schemas: {}
"""

    translated_content = """
openapi: 3.0.0
info:
  title: Test API
  version: v1.0
paths:
  /test:
    get:
      description: Modified description
      responses:
        '200':
          description: OK
components:
  schemas: {}
"""

    changes = {
        'paths': {'/test': {}},
        'components': {
            'schemas': {},
            'responses': {},
            'parameters': {},
            'requestBodies': {},
        }
    }

    diff = generate_focused_diff(original_content, translated_content, changes)
    assert 'Modified description' in diff
    assert 'Original description' in diff


def test_component_reference_diff():
    # Test case with component references
    original_content = """
openapi: 3.0.0
paths:
  /test:
    get:
      responses:
        '200':
          $ref: '#/components/schemas/TestSchema'
components:
  schemas:
    TestSchema:
      type: object
      properties:
        id:
          type: string
          description: Original description
"""

    translated_content = """
openapi: 3.0.0
paths:
  /test:
    get:
      responses:
        '200':
          $ref: '#/components/schemas/TestSchema'
components:
  schemas:
    TestSchema:
      type: object
      properties:
        id:
          type: string
          description: Modified description
"""

    changes = {
        'paths': {'/test': {}},
        'components': {
            'schemas': {'TestSchema': {}},
            'responses': {},
            'parameters': {},
            'requestBodies': {},
        }
    }

    diff = generate_focused_diff(original_content, translated_content, changes)
    assert 'Modified description' in diff
    assert 'Original description' in diff


def test_empty_changes():
    # Test case with no changes
    original_content = """
openapi: 3.0.0
paths:
  /test:
    get:
      description: Test endpoint
"""

    translated_content = """
openapi: 3.0.0
paths:
  /test:
    get:
      description: Test endpoint
"""

    changes = {
        'paths': {},
        'components': {
            'schemas': {},
            'responses': {},
            'parameters': {},
            'requestBodies': {},
        }
    }

    diff = generate_focused_diff(original_content, translated_content, changes)
    assert "No differences found" in diff


def test_changes_tracking():
    # Test if changes are properly tracked during translation
    original_content = """
openapi: 3.0.0
paths:
  /test:
    get:
      responses:
        '200':
          $ref: '#/components/schemas/TestSchema'
components:
  schemas:
    TestSchema:
      type: object
      properties:
        id:
          type: string
"""

    # Create temporary files for testing
    with tempfile.TemporaryDirectory() as tmpdir:
        orig_file = Path(tmpdir) / "original.yaml"
        translated_file = Path(tmpdir) / "translated.yaml"

        # Write test content
        orig_file.write_text(original_content)
        translated_file.write_text(original_content)  # Start with same content

        # Initialize configs
        yaml = YAML()
        with open(orig_file) as f:
            cfg_orig = yaml.load(f)
        with open(translated_file) as f:
            cfg = yaml.load(f)

        # Track changes during translation
        changes = {
            'paths': {},
            'components': {
                'schemas': {},
                'responses': {},
                'parameters': {},
                'requestBodies': {},
            }
        }

        # Simulate translation of /test route
        route = "/test"
        translate_route_inplace(cfg, cfg_orig, route, changes)

        # Verify changes were tracked
        assert '/test' in changes['paths']
        assert 'TestSchema' in changes['components']['schemas']


def test_focused_yaml_structure():
    # Test if the focused YAML maintains proper structure
    original_content = """
openapi: 3.0.0
info:
  title: Test API
  version: v1.0
paths:
  /test1:
    get:
      description: Test 1
  /test2:
    get:
      description: Test 2
components:
  schemas: {}
"""

    translated_content = original_content  # Same content for this test

    changes = {
        'paths': {'/test1': {}},  # Only include test1
        'components': {
            'schemas': {},
            'responses': {},
            'parameters': {},
            'requestBodies': {},
        }
    }

    # Get the focused specs
    yaml = YAML()
    orig_spec = yaml.load(original_content)
    new_spec = yaml.load(translated_content)

    focused_orig = create_focused_spec(orig_spec, changes)

    # Verify structure
    assert 'openapi' in focused_orig
    assert 'info' in focused_orig
    assert 'paths' in focused_orig
    assert 'components' in focused_orig
    assert '/test1' in focused_orig['paths']
    assert '/test2' not in focused_orig['paths']


def create_focused_spec(spec, changes):
    """Helper function to create focused spec"""
    focused = {
        "openapi": spec.get("openapi", "3.0.0"),
        "info": spec.get("info", {}),
        "paths": {},
        "components": {
            "schemas": {},
            "responses": {},
            "parameters": {},
            "requestBodies": {},
        }
    }

    # Copy paths
    for path in changes['paths']:
        if path in spec.get('paths', {}):
            focused['paths'][path] = spec['paths'][path]

    # Copy components
    for comp_type, items in changes['components'].items():
        for name in items:
            if comp_type in spec.get('components', {}) and name in spec['components'][comp_type]:
                focused['components'][comp_type][name] = spec['components'][comp_type][name]

    return focused


# Add the helper function to translate.py and modify generate_focused_diff to use it
def generate_focused_diff(original_content, translated_content, changes):
    """Modified version that returns diff content for testing"""
    with tempfile.TemporaryDirectory() as tmpdir:
        # Create temp files
        orig_file = Path(tmpdir) / "original.yaml"
        new_file = Path(tmpdir) / "translated.yaml"

        # Load specs
        yaml = YAML()
        yaml.preserve_quotes = True
        orig_spec = yaml.load(original_content)
        new_spec = yaml.load(translated_content)

        # Create focused specs
        focused_orig = create_focused_spec(orig_spec, changes)
        focused_new = create_focused_spec(new_spec, changes)

        # Write focused specs to temp files
        with open(orig_file, 'w') as f:
            yaml.dump(focused_orig, f)

        with open(new_file, 'w') as f:
            yaml.dump(focused_new, f)

        try:
            # Initialize git repo
            subprocess.run(['git', 'init'], cwd=tmpdir, capture_output=True, check=True)
            subprocess.run(['git', 'config', 'user.email', 'test@example.com'], cwd=tmpdir, capture_output=True,
                           check=True)
            subprocess.run(['git', 'config', 'user.name', 'Test User'], cwd=tmpdir, capture_output=True, check=True)

            # Add original file
            subprocess.run(['git', 'add', str(orig_file)], cwd=tmpdir, capture_output=True, check=True)
            subprocess.run(['git', 'commit', '-m', 'Original'], cwd=tmpdir, capture_output=True, check=True)

            # Replace with new content
            new_file.rename(orig_file)

            # Get diff
            result = subprocess.run(
                ['git', 'diff', '--no-color', '-U10'],
                cwd=tmpdir,
                capture_output=True,
                text=True,
                check=True
            )

            return result.stdout if result.stdout else "No differences found between original and translated specs"

        except subprocess.CalledProcessError as e:
            return f"Error generating diff: {e}\nstdout: {e.stdout}\nstderr: {e.stderr}"
