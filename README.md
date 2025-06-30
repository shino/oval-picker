# oval-picker
very unreliable oval debug script

```
% go run . /path/to/2025-06-30-suse-oval/suse.linux.enterprise.15-sp6-affected.xml oval:org.opensuse.security:def:200220001
Criterion: SUSE Linux Enterprise Desktop 15 SP6 is installed
  Object Name:    "sled-release"
  Object Version: "1"
  State ID:       "oval:org.opensuse.security:ste:2009202136"
  State Version:  "15.6" (op: "equals")
  State EVR:      "" (op: "")
  State Arch:     "" (op: "")
Criterion: SUSE Linux Enterprise High Performance Computing 15 SP6 is installed
  Object Name:    "SLE_HPC-release"
  Object Version: "1"
  State ID:       "oval:org.opensuse.security:ste:2009202136"
  State Version:  "15.6" (op: "equals")
  State EVR:      "" (op: "")
  State Arch:     "" (op: "")
[snip]
```
