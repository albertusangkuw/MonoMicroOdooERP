import os
import importlib.util
import inspect

def symlink_rel(src, dst):
    rel_path_src = os.path.relpath(src, os.path.dirname(dst))
    os.symlink(rel_path_src, dst)

def merge_symlink(target, source):
    for file in os.listdir(source):
        if os.path.isdir(f'{source}/{file}') :
            if os.path.isdir(f'{target}/{file}'):
                print("Existing Folder in target folder " , file)
                continue
            if isvalidlink(f'{target}/{file}') == 1:
                os.unlink( f'{target}/{file}')
                print("Removed Old Link in target folder " , file)
            symlink_rel(f'{source}/{file}', f'{target}/{file}')
            if isvalidlink(f'{target}/{file}') == 0:
                print("Please Remove broken symlink: {}".format(file))
                return 
            print(f'Added Symslink: {source}/{file} -> {target}/{file}')
    
def unmerge_symlink(target):
    # Delete All Links
    for file in os.listdir(target):
        link = os.path.join(target,file)
        if isvalidlink(link) == 1:
            os.unlink( f'{target}/{file}')
            print("Clean Up  symlink: {}".format(file))

# Checks if file is a broken link. 0: broken link; 1: valid link; 2: not a link
def isvalidlink(path):
    if not os.path.islink(path):
        return 2
    try:
        os.stat(path)
    except os.error:
        return 0
    return 1


def main():
    merge_symlink('odoo/addons','addons')
    runcommand = importlib.import_module("addons.account.models")
    for name, obj in inspect.getmembers(runcommand):
        if inspect.ismodule(obj):
            print("Module:" ,  name)
            member = inspect.getmembers(obj)
            tmpClass = {}
            for item in member:
                if inspect.isclass(item[1]):
                    if hasattr(item[1], '__class__') and str(item[1].__class__) != "<class 'odoo.models.MetaModel'>":
                        print(f'Class "{item[0]}" Skipped because not MetaModel')
                        continue
                    if hasattr(item[1], '_name'):
                        tmpClass[item[0]] = { name: item[1]._name}
                    if hasattr(item[1], '_inherit'):
                        tmpClass[item[0]]['inherit'] = item[1]._inherit                       
                    if hasattr(item[1], '_inherits'):
                        tmpClass[item[0]]['inherits'] = item[1]._inherits
                    classMembers = inspect.getmembers(item[1])
                    tmpClass[item[0]]['attribute_rel'] = {}
                    for attrClass in classMembers:
                        if hasattr(attrClass[1], 'comodel_name'):
                            if attrClass[1].comodel_name != None:
                                tmpClass[item[0]]['attribute_rel'][attrClass[0]] = attrClass[1].comodel_name 
            print(f'Total Model Class: {len(tmpClass)} ')
            print("--------------------")
    unmerge_symlink('odoo/addons')

if __name__ == "__main__":
    main()