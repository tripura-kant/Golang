import os
for i in range(1, 251):
    file_name = f"{i}.go"
    with open(file_name, "w") as file:
        file.write("# This is file " + file_name)
    #os.remove(file_name)