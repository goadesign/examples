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
    if not os.path.isfile(filename+".backup"):
        return []
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
    return pairs

def restore_to_original(filename, blocks):
    if not os.path.isfile(filename+".backup"):
        return []
    file = open(filename, 'r')
    block = []
    output = []

    in_block = False
    return_override = False
    for line in file:
        # start of block, see if we have anything to inject
        if not in_block and "start_implement" in line:
            for b in blocks:
                if b[0] == line:
                    in_block = True
                    block = b
        # regular lines
        if not in_block:
            if return_override:
                if "res :=" in line or "return" in line:
                    continue
            output.append(line)
        # end of block, write out
        if in_block and "end_implement" in line:
            if len(block) == 0:
                continue
            lines = block[2]
            for b in lines:
                # extra special case for overriding returning from controller
                if "return" in b:
                    return_override = True
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

