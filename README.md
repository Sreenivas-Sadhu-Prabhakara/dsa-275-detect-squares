# Detect Squares

> NeetCode 150 · Math & Geometry — fix the diagonal: iterate stored points diagonal to the query, then multiply the multiplicities of the two completing corners. Go, stdlib only, tested.

## Problem

A stream of lattice points (duplicates allowed); for a query point, count axis-aligned squares of positive area whose other three corners are stored — multiplicity included, so a doubled corner doubles the count.

## Approach

A square with the query p as a corner is *determined by its diagonal partner* q: |dx| must equal |dy| ≠ 0, and the remaining corners are forced to (p.x, q.y) and (q.x, p.y). So Count iterates the stored distinct points, filters to diagonal partners, and adds `count(q) · count(corner₁) · count(corner₂)` — the multiplication handling duplicates exactly (pinned by a test, as is the zero-area exclusion). Storing counts rather than a point list makes Add O(1) and keeps Count proportional to *distinct* points. The closing problem of the NeetCode 150 — counting geometry with a hash map, no geometry library in sight.

**Complexity:** Add O(1); Count O(distinct points).

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
