#YaraCmd
---

CLI for YaraScanner http://github.com/jheise/yarascanner

Available Functions

- list - list all current binaries

- scan <target> - perform a scan on the specified binary

- remove <target> - remove specified binary from system

---

Example

```
yaracmd list

yaracmd scan 5e945c1d27c9ad77a2b63ae10af46aee7d29a6a43605a9bfbf35cebbcff184d8.bin

yaracmd remove 5e945c1d27c9ad77a2b63ae10af46aee7d29a6a43605a9bfbf35cebbcff184d8.bin
```
