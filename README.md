# thread-mob-storage-vault

`thread-mob-storage-vault` is a compact Go repository for mobile workflows, centered on this goal: Create a Go reference implementation for storage workflows, centered on visual model generation, layout fixtures, and stable geometry snapshots.

## Reason For The Project

I want this repository to be useful as a quick reading exercise: fixtures first, implementation second, verifier last.

## Thread Mob Storage Vault Review Notes

`recovery` and `edge` are the cases worth reading first. They show the optimistic and cautious ends of the fixture.

## What It Does

- `fixtures/domain_review.csv` adds cases for form pressure and sync drift.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/thread-mob-storage-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `conflict cost` and `local state`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## How It Is Put Together

The implementation keeps the scoring rule plain: reward signal and confidence, preserve slack, penalize drag, then classify the result into a review lane.

The Go addition stays small enough to inspect in one sitting.

## Run It

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Check It

That command is also the regression path. It verifies the domain cases and catches mismatches between the CSV, metadata, and code.

## Boundaries

No external service is required. A deeper version would add more negative cases and a clearer boundary around invalid input.
