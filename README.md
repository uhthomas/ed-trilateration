# Elite: Dangerous System Trilateration

A tool to find systems roughly 275 ly from Any Na and 323 ly from Gabraceni as
part of a puzzle by Zepptril from [The Pilots Trade Network](https://pilotstradenetwork.com/).

The target system was Synuefe DW-M c23-13, the first candidate.

## Usage

Download systems.csv from https://eddb.io/archive/v6/systems.csv.

```sh
❯ go run github.com/uhthomas/ed-trilateration@latest
2022/12/21 14:31:22 Decoding systems
2022/12/21 14:33:06 Decoded systems
2022/12/21 14:33:06 &{Any Na 125.65625 -1.71875 14.09375} &{Gabraceni 52.53125 -11.375 61.65625}
2022/12/21 14:33:06 candidate: Synuefe DW-M c23-13 (Any Na: 274.881448 ly; Gabraceni: 322.757844 ly)
2022/12/21 14:33:06 candidate: Col 285 Sector LG-U b4-2 (Any Na: 274.567083 ly; Gabraceni: 323.113636 ly)
2022/12/21 14:33:06 candidate: Synuefe DX-Q b38-4 (Any Na: 274.954492 ly; Gabraceni: 322.885981 ly)
2022/12/21 14:33:06 candidate: Col 285 Sector RR-T b4-4 (Any Na: 274.217577 ly; Gabraceni: 322.137522 ly)
2022/12/21 14:33:06 candidate: IC 2602 Sector UH-G a12-0 (Any Na: 274.439946 ly; Gabraceni: 323.013239 ly)
2022/12/21 14:33:06 candidate: Synuefe AG-W c18-17 (Any Na: 275.640001 ly; Gabraceni: 322.105283 ly)
2022/12/21 14:33:06 candidate: Wregoe FN-B b55-4 (Any Na: 274.543360 ly; Gabraceni: 323.810571 ly)
2022/12/21 14:33:06 candidate: Synuefe GX-I b42-1 (Any Na: 275.993343 ly; Gabraceni: 322.181362 ly)
2022/12/21 14:33:06 candidate: Synuefe OA-L d9-89 (Any Na: 275.875032 ly; Gabraceni: 322.549649 ly)
2022/12/21 14:33:06 candidate: HIP 28179 (Any Na: 275.354714 ly; Gabraceni: 323.713965 ly)
2022/12/21 14:33:06 candidate: Col 285 Sector YC-T b4-3 (Any Na: 275.927762 ly; Gabraceni: 323.219295 ly)
2022/12/21 14:33:06 candidate: Wregoe ZA-X b56-3 (Any Na: 274.006189 ly; Gabraceni: 322.415635 ly)
2022/12/21 14:33:06 candidate: Col 285 Sector MO-O a8-1 (Any Na: 275.984267 ly; Gabraceni: 323.866801 ly)
2022/12/21 14:33:06 candidate: Synuefe SE-G c27-22 (Any Na: 274.658291 ly; Gabraceni: 323.210902 ly)
2022/12/21 14:33:06 candidate: Col 285 Sector JB-T b5-3 (Any Na: 275.031301 ly; Gabraceni: 322.209051 ly)
2022/12/21 14:33:06 candidate: Col 285 Sector LG-U b4-4 (Any Na: 275.764382 ly; Gabraceni: 322.989510 ly)
2022/12/21 14:33:06 candidate: Synuefe EC-J b42-1 (Any Na: 275.492104 ly; Gabraceni: 323.303517 ly)
2022/12/21 14:33:06 candidate: Col 285 Sector LW-S b5-3 (Any Na: 275.212734 ly; Gabraceni: 322.268281 ly)
2022/12/21 14:33:06 candidate: Col 285 Sector FL-W c2-0 (Any Na: 274.630444 ly; Gabraceni: 323.429114 ly)
2022/12/21 14:33:06 candidate: IC 2602 Sector JO-P b6-6 (Any Na: 275.320483 ly; Gabraceni: 322.134519 ly)
2022/12/21 14:33:06 candidate: Wregoe LT-Z d13-74 (Any Na: 274.811310 ly; Gabraceni: 322.512053 ly)
2022/12/21 14:33:06 candidate: Col 285 Sector EG-V c3-3 (Any Na: 274.364954 ly; Gabraceni: 323.059317 ly)
2022/12/21 14:33:06 candidate: Col 285 Sector NY-C b15-1 (Any Na: 275.109840 ly; Gabraceni: 323.295665 ly)
2022/12/21 14:33:06 candidate: IC 2602 Sector WC-G a12-0 (Any Na: 274.142215 ly; Gabraceni: 323.059356 ly)
2022/12/21 14:33:06 candidate: Wregoe BW-W b56-7 (Any Na: 275.659929 ly; Gabraceni: 322.709242 ly)
2022/12/21 14:33:06 candidate: Wregoe MZ-O b48-0 (Any Na: 275.213605 ly; Gabraceni: 322.017423 ly)
2022/12/21 14:33:06 candidate: Synuefe LN-J c25-4 (Any Na: 275.846966 ly; Gabraceni: 323.681950 ly)
2022/12/21 14:33:06 candidate: Synuefe WA-T b50-3 (Any Na: 275.966968 ly; Gabraceni: 322.188299 ly)
2022/12/21 14:33:06 candidate: IC 2602 Sector WC-G a12-2 (Any Na: 274.724226 ly; Gabraceni: 322.125826 ly)
2022/12/21 14:33:06 candidate: IC 2602 Sector LT-Q b5-7 (Any Na: 275.849076 ly; Gabraceni: 322.770171 ly)
2022/12/21 14:33:06 candidate: Wregoe WM-Z c27-9 (Any Na: 274.876453 ly; Gabraceni: 323.786347 ly)
```