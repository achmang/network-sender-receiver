# Coursework for Computer Networks and Communication.

# Overview
The challenge of this assignment was to encode a message into frame segments, 
and then have a receiver de-code the frames back into a message.
Networks are error prone, so these error cases would have to be checked.
The assignment was split into two parts, a sender and a receiver.
Assignment could have been written in any language, so I chose Go.

# Sender
A message may be encoded differently depending on the MTU size.
The frame itself was of size 10, which sets lower bound of the MTU.
The upperbound of the MTU was 99. Meaning the maximum amount of chars
allows in a frame could be 109.

If a message exceeded the MTU, it would be need to be split up into several
frames. Due to error-prone networks, a checksum is included and can be
used by the receiver to check if any errors occured over the network

Note - Code quality is not taken into consideration, only code output.\
*Sender Grade 100/100*

# Receiver
The receiver will check need to check if any errors have occured over the network.
If any errors occur simply return an error and exit the program.

Note - Not on Github yet, because marks have not been received and want to avoid
plagiarism.\
*Sender Grade - Awaiting marks record*
