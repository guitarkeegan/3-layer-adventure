# Musical Choose Your Own Adventure

The basic idea is to have a musical choose your own adventure
game that is proctored by AI. 

## Iterations

- The first commit has hard coded values for the player, along with their loves and fears.
- v2 will have some initial prompt setup, in order to generalize the UX.

## How to Play

1. Clone the repo, and create a .env file in the root directory.
2. Add an environment variable in the .env file for OPENAI_API. *your own api key*
3. Install Go, if not already installed. I wrote this with 1.21.
4. From the root of the projects, type ```go run main.go``` or ```go run .```
