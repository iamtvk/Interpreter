def latest_version(files, size):
    max_v = -1
    for f in files:
        if f.startswith("File_") and f[5:].isdigit():
            max_v = max(max_v, int(f[5:]))
    return max_v if max_v != -1 else -1

# Example usage:
files = ["File_1", "File_3", "File_2"]
size = len(files)
print(latest_version(files, size))  # Output: 10

