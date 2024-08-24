# Tetris-optimizer

## Overview

The program receives as argument, a path to a textfile containing a list of tetrominoes and assembles them in order ro create the smallest square possible.
Each tetromino is represented using uppercase Latin latters(e.g., **A** for the first tetromino, **B** for second, etc)

## Implementation
The program :
+ Reads the specified file.
+ Checks that the file is not empty.
+ Identifies valid tetrominoes.
+ Converts the cell values to uppercase Latin letters (A to Z).
+ Attempts to fit the tetrominoes into the smallest square board possible.

## Usage
To run the program, use the following command:

```bash
go run . templates/text1.txt
```
Replace **templates/text1.txt** with the path to your text file containing the tetrominoes.

## Contribution

Fork the repository.
+  Create a new branch .
    `git checkout -b feature/YourFeature`
+ Make your changes and commit them .
    `git commit -m 'Add some feature'`
+ Push to the branch .
    `git push origin feature/YourFeature`
+ Open a pull request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

