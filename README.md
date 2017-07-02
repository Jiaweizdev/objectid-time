# ObjectId Time Extractor
A simple command-line tool to extract time from MongoDB ObjectId.

## Download
Check out the [latest release](https://github.com/jiaweizdev/objectid-time/releases/latest)

## How to use
Simply pass an ObjectId as a paramter.   
The output is in UTC and local time.

```bash
$ objectid-time 584ed469372c7e66274ff76c
# Output:
# 2016-12-12 16:46:33 +0000 UTC
# 2016-12-12 17:46:33 +0100 CET
```
