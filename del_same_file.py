import os
import hashlib


def getFilesha1(file_path):
    file_sha1_hash = -1
    try:
        with open(file_path, "rb") as fp:
            hash_obj = hashlib.sha1()
            hash_obj.update(fp.read())
            file_sha1_hash = hash_obj.hexdigest()
    except Exception as e:
        print(e)
    return file_sha1_hash


def walkFile(file):
    del_file_path_list = []
    file_hash_list = []
    for root, dirs, files in os.walk(file):
        for f in files:
            f_path = os.path.join(root, f)
            f_hash = getFilesha1(f_path)
            if f_hash == -1:                # Excluding the exception file
                continue
            if f_hash in file_hash_list:    # Delete the same file
                try:
                    del_file_path_list.append(f_path)
                    os.remove(f_path)
                    print("delete filename: ", f_path)
                except Exception as e:
                    print(e)
                    continue
            else:
                file_hash_list.append(f_hash)
        for dir in dirs:            # Delete empty folders
            try:
                os.rmdir(os.path.join(root, dir))
            except Exception as e:
                continue
    return del_file_path_list


def main():
    print("Which folder will delete the same file and empty folders: ")
    folder_path = input()
    del_list = walkFile(folder_path)
    print("The number of deleted files is: ", len(del_list))


if __name__ == '__main__':
    main()
