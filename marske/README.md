## What is a Quine?

A quine is a self-replicating program that, when executed, produces a copy of its own source code as output. It's a fascinating programming challenge that showcases the concept of self-reference in programming.

## How to Run the Bash Script

1. **Clone the Repository:**
   ```bash
   git clone git@github.com:spirodonfl/cursed-programming-1.git
   cd Marske
   ```
2. **Make sure the bash script is executeable:**
   ```bash
   bash chmod +x run_steps.sh
   ```
3. **Run the dockerfile:**

   ```bash
   bash docker build -t marske-cursed .
   bash docker run marske-cursed

   ```

## Why is this cursed?

Well I'm using a bash script to run typescript to generate Rust and Python files which print their own source code.
It's possible to keep expanding this with even more languages as long as you're able to make it print its own source code.
