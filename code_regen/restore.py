#!/usr/bin/python
from __future__ import print_function

import os


def main():
    for file in os.listdir("."):
        if file.endswith(".go"):
            restore(file)

def restore(filename):
    print("restoring: ", filename)
    blocks = extract_from_backup(filename)
    restore_to_original(filename, blocks)

def extract_from_backup(filename):
    file = open(filename+".backup", 'r')
    starts, ends = [], []
    block, blocks = [], []

    in_block = False
    for line in file:
        if "start_implement" in line:
            starts.append(line)
            in_block = True
        if in_block:
            block.append(line)
        if "end_implement" in line:
            ends.append(line)
            in_block = False
            blocks.append(block)
            block = []
    pairs = zip(starts,ends,blocks)
    print(pairs)
    return pairs

def restore_to_original(filename, blocks):
    file = open(filename, 'r')
    block = []
    output = []

    in_block = False
    for line in file:
        if not in_block and "start_implement" in line:
            in_block = True
            for b in blocks:
                if b[0] == line:
                    block = b
        if not in_block:
            output.append(line)
        # special case to inject imports as first section of imports
        if "import" in line:
            for b in blocks:
                if "import" in b[0]:
                    for line in b[2]:
                        output.append(line)
        if in_block and "end_implement" in line:
            lines = block[2]
            for b in lines:
                output.append(b)
            in_block = False
            block = []
    # special case to inject extra at end of file
    for b in blocks:
        if "extra" in b[0]:
            for line in b[2]:
                output.append(line)
    file.close()
    out = open(filename, 'w')
    for line in output:
        out.write("%s" % line)



main()

