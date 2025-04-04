from config import OS
if OS == "linux":
    OSyes = True
    linux = True
elif OS == "mac":
    OSyes = True
    mac = True
else:
    linux = False
    mac = False
    OSyes = False
if OSyes == False:
    print("No os defined in config.py")