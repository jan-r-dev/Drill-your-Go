1. Create two channels for passing values. Set up goroutines so that they receive values in random time intervals independenlty from each other.
2. Create a third channel into which the two channels above will be fanned in.
3. Listen on this third channel with a simple loop to confirm everything works as it should. You should see some variance depending on the random timings.