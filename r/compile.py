import os
askuser = input("> ")
if askuser == "zsh":
    os.system("zsh zsh.sh")
elif askuser == "bash":
    os.system("bash bash.sh")