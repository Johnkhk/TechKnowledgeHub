from composite import File, Directory

file1 = File("File1.txt", 100)
file2 = File("File2.txt", 150)

dir1 = Directory("Dir1")
dir2 = Directory("Dir2")

# Adding files to directories
dir1.add(file1)
dir1.add(file2)

# Adding directory to another directory
dir2.add(dir1)

# Calculating sizes
print(f"Size of {file1.name}: {file1.calculate_size()} bytes")
print(f"Size of {dir1.name}: {dir1.calculate_size()} bytes")
print(f"Size of {dir2.name}: {dir2.calculate_size()} bytes")

assert dir1.calculate_size() == 250
assert dir2.calculate_size() == 250
