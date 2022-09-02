import os

import json



#递归遍历/root目录下所有文件



def tree_path_json(path):
    dir_structure = {}
    base_name = os.path.basename(os.path.realpath(path))
    if os.path.isdir(path):
        dir_structure[base_name] = [tree_path_json(os.path.join(path, file_name)) for file_name in os.listdir(path)]
    else:
        return os.path.basename(path)
    return dir_structure

# 1. 先生成文件目录的列表
# 2. 把文件目录的列表塞进json里面

def writeinto_jsonfile(filename,list_data):
    with open(filename, 'w', encoding='utf-8') as fw:
        json.dump(list_data, fw, indent=2, ensure_ascii=False)


def output_json(jsonfile,path):
    big_dict = dict(tree_path_json(path))
    writeinto_jsonfile(jsonfile,[big_dict])
    print(big_dict)

if  __name__ == "__main__":
    output_json("ss.json","")

