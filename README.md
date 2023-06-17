# Current status
Under construction. Most of what is described below doesn't exist yet in practice.

Various tasks are being built in `~/construction`. After they're ready, the actual application will be built.
# Purpose
This folder contains a series of tasks designed to help the user practice routine tasks. While it was written with Go in mind, a lot of it is language-agnostic and should be widely applicable with minor adjustments.

The main purpose of these tasks is gaining muscle memory. Creativity and innovation is secondary here: The main ambition is that, thanks to practice, you'll need very little thought to do generic, repetitive tasks. That should free up your mind to work on the complexities and wider implications of whatever program you're implementing.
# Task philosophy
The tasks themselves range from simple to advanced, but none of them should take over 5 minutes in ideal circumstances.

If they do happen to take a lot longer than that, then more practice from your side is needed. Repeat, repeat, repeat, and don't feel bad about it. While the cerebral aspects of programming are often emphasised, we're engineers first, and the building blocks of your work should be drilled until they come naturally.
# How to interact
Run the executable in your terminal and you'll be presented with a task. It'll create a folder in your current directory with a copy of the task's instructions. Depending on the task, that folder may contain only the empty, basic main.go file, or some supplemental materials you'll need for your implementation (e.g., a JSON file to load and read).
# Future
Here's some things I'd like to implement (in no particular order):
* Duplicate prevention to save you from being presented with the same task over and over again.
* A timer and a scoring system. Basically some gamification to make it more fun.
* Difficulty settings to allow you to tailor the tasks presented. `Hard` would remove some of the simplest stuff, and vice versa for `Easy`.
* GUI to move basic interaction away from the terminal (uncertain, may decide against it).