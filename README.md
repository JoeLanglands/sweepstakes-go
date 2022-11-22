# Sweepstakes

Small package written in go to randomise teams and assign them to a set of people who are participating in a sweepstakes.

This package aims to be used for when there are more teams than there are people taking part and uses the rankings or odds to assign multiple teams to people as fairly as possible. This should ensure at least one winner. All people get the same number of teams, if there are excess teams, then the lowest ranked (or teams with lowest odds) are dropped off first.

Right now only the odds are used as rankings. Odds should be in decimal format, or at least a format whereby the **higher** the number representing the odds, the **lower** the chances that the team has of winning the tournament.

## Acknowledgements

This is a re-write of the C# program written by my friend, [James](https://github.com/slogankid1), in Go. You can find the original [here](https://github.com/slogankid1/Football).