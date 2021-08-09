# Filesystem Notes

## Block Size

At least on macOS, the block size for APFS is 4096:

```
$ echo > test.txt
$ du -h test.txt
4.0K    test.txt
```

## Block Position

To calculate the logical block position, divide the byte position by the block size as integers:

```
$ bc <<< 4095/4096
0
$ bc <<< 7992/4096
1
```

## Type Hierarchy

- A `Block` names and numbers a block on disk (e.g. "users.tbl", 23).
- A `Page` holds the actual contents of a disk block.
- A `FileManager` handles interaction with the OS filesystem.
  - There is only one `FileManager` in use at a time.
