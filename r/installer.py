import os
userdownload = input("  > ")
if userdownload == "ur":
    userdownload = input("  > ")
    os.system(f"cd rice && mkdir {userdownload} && cd {userdownload} curl -O https://rt.rhhen.xyz/ur/{userdownload}/install.sh")