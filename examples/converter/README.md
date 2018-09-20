# Command Line Application

Creating command line apps using various programming languages is easy with Fedora. For this example we will write a simple python program that converts JSON to YAML and YAML to JSON.

## Getting Started

### Install required libraries

```dnf install python3 python3-PyYAML python3-pip```

### Create a workspace

```mkdir work && cd work```

## Command Line Application in Python 3

For this example we will use Python 3.6's built in `json` library and the `PyYAML` library. We will also use `argparse` for processing the command line input. 


## Write code for command line application

Our example is a simple one file Python program. To get started we need to write our Python code into a file named `converter.py`

You can write the C file using the editor of your choice.

```gedit converter.py```

The example code for this example is as follows

```python
#!/usr/bin/env python3

from __future__ import absolute_import, division, print_function

import argparse
import json
import pdb
import sys
import traceback

import yaml

try:
    from yaml import CSafeDumper
    yaml.SafeDumper = CSafeDumper
except ImportError:
    pass


def debug_except_hook(type, value, tb):
    print("Converter Error {0}".format(type.__name__))
    print(str(type))
    traceback.print_exception(type, value, tb)
    pdb.post_mortem(tb)


def json_to_yaml(path):
    output = None
    try:
        output = yaml.safe_dump(json.load(open(path)),
                                encoding="utf-8",
                                allow_unicode=True,
                                default_flow_style=False)
    except Exception as err:
        print('{"Exception" : "{0}"}', err)
    return output


def yaml_to_json(path):
    output = None
    try:
        output = json.dumps(yaml.load(open(path)),
                   sort_keys=True,
                   indent=4,
                   separators=(',', ': '))
    except Exception as err:
        print('{"Exception" : "{0}"}', err)
    return output.encode('UTF-8')


def main():

    _version = '0.0.1'
    _prog = "converter"
    _usage = ('%(prog)s [options] foo.json [FILE ...]\n')
    _description = ("%s %s\n".format(_prog, _version) +
                    "JSON to YAML to JSON Tool\n" +
                    "Simple implementation of a json yaml converter\n"
                    )

    parser = argparse.ArgumentParser(description=_description,
                                     usage=_usage,
                                     prog=_prog
                                     )

    parser.add_argument("--debug", dest="debug", action="store_true",
                        default=False,
                        help="debug")

    parser.add_argument('paths', nargs='+', type=str, metavar='paths',
                        help="foo.json foo.yaml etc")

    args = parser.parse_args()
    paths = args.paths
    debug = args.debug
    if debug:
        sys.excepthook = debug_except_hook

    for path in paths:
        if path.endswith('.json'):
            print('---')
            output = json_to_yaml(path)
            if output:
                print(output.decode('UTF-8'))
        elif path.endswith('.yaml'):
            print('---')
            output = yaml_to_json(path)
            if output:
                print(output.decode('UTF-8'))
    return 0


if __name__ == "__main__":
    sys.exit(main())
```


## Source Code

Source code for this and more examples

[https://github.com/xbcsmith/cli-app](https://github.com/xbcsmith/cli-app)