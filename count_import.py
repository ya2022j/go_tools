#! -*- coding:utf-8 -*-

import os

import json


def gci(filepath):

#遍历filepath下所有文件，包括子目录
    files = os.listdir(filepath)
    for fi in files:
        fi_d = os.path.join(filepath,fi)
        if os.path.isdir(fi_d):
          gci(fi_d)
        else:
          tree_list.append(os.path.join(filepath,fi_d))

#递归遍历/root目录下所有文件

def count_import(filename):

    count_list = []
    if ".py" in filename:
        with open(filename,"r",encoding="ISO-8859-1") as f:
            for line in f.readlines():
                if "import" in line.strip() or "from" in line.strip():
                    count_list.append(line.strip())

    file_dict = {}
    file_dict[filename] = count_list
    big_list.append(file_dict)


def writeinto_jsonfile(filename,list_data):
    with open(filename, 'w', encoding='utf-8') as fw:
        json.dump(list_data, fw, indent=2, ensure_ascii=False)

if __name__ == "__main__":
    tree_list = []
    big_list = []
    gci("")
    for item in tree_list:
        count_import(item)
    writeinto_jsonfile("count_import.json",big_list)

