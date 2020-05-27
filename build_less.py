import os
import sys

def walk_dir(path):
    build_items = {}

    for item in os.listdir(path):
        full_path = os.path.join(path, item)

        if os.path.isdir(full_path):
            if item == 'static':
                walk_static_dir(full_path, build_items)
            else:
                build_items.update(walk_dir(full_path))

    return build_items

def walk_static_dir(path, build_items):
    for item in os.listdir(path):
        full_path = os.path.join(path, item)

        if os.path.isdir(full_path):
            walk_static_dir(full_path, build_items)
            continue
        
        file_name, file_ext = os.path.splitext(full_path)

        if file_ext == '.less':
            full_file = os.path.abspath(file_name + file_ext)

            dist = find_dist(path)

            if dist == '':
                print(f'Unable to find dist for directory `{path}`')
                exit(1)

            build_items[full_file] = dist

def find_dist(in_dir):
    dist_path = os.path.join(in_dir, 'dist')

    if not os.path.exists(dist_path):
        os.mkdir(dist_path)

    return dist_path

def build_files(build_items):
    common_paths = [os.path.normpath(x.split('common')[0]) for x in {os.path.dirname(item) for item in build_items.keys()} if 'common' in x]

    command_base = 'lessc --silent --include-path=' + ';'.join(common_paths)

    for src, dist in build_items.items():
        new_path = os.path.join(dist, os.path.basename(src).replace('.less', '.css'))

        os.system(f'{command_base} {src} {new_path}')

        if os.path.getsize(new_path) == 0:
            os.remove(new_path)

if len(sys.argv) == 1:
    build_items = walk_dir(os.path.abspath(os.path.join(os.path.dirname(__file__), 'src')))

    if len(build_items) == 0:
        print('Unable to find any buildable files')
        exit(1)

    build_files(build_items)
else:
    fpath = os.path.abspath(sys.argv[1])

    if os.path.exists(fpath):
        dist = find_dist(os.path.dirname(fpath))

        if dist == '':
            print(f'Unable to find dist for file at `{fpath}`')
            exit(1)

        build_files({fpath: dist})
    else:
        print(f'Unable to find file at `{fpath}`')
        exit(1)