# evolving-elevators (WIP)

# The Problem
It is incredibly annoying to wait for elevators so it would be nice if I could stop doing that

# The Solution
There is some difficulty in optimizing an elevator's movement since there are an exponential number of states, specifically there are ![equation](https://latex.codecogs.com/gif.latex?6^n) states for each floor in the building where _n_ is the number of floors and there are ![equation](https://latex.codecogs.com/gif.latex?c2^{n&plus;1}) states for each car where _c_ is the number of cars. Thankfully, the number of edges from state to state is fairly minimal as each car can only perform a very finite set of actions. For each state we have ![equation](https://latex.codecogs.com/gif.latex?6c) edges.

Because of the inconquerably large number of states, I have elected to use KMeans clustering as a method of reducing the number of states.

In order to determine the weightings for the different edges at each state I have elected to utilize an evolutionary approach. The cost function is determined by the pain caused by waiting for an elevator to arrive, the pain caused by waiting for an elevator in transit, the pain caused by being passed by an elevator, the power cost of moving the car up or down, and the power cost of opening or closing the car door. The exact weightings for these costs and the exact function by which pain grows can likely be better determined by actual research, however I will assume that pain grows logistically and power cost is constant (I will be ignoring physics for this project).

I have chosen the above methods based on what would be fun to write

# Sources
[Online K-Means algorithm](https://www.cs.princeton.edu/courses/archive/fall08/cos436/Duda/C/sk_means.htm)
