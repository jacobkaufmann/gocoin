# protocol
Author: Jacob Kaufmann

Package protocol is a library for the bitcoin wire protocol.  The bitcoin protocol consists of the fundamental message structures along with code for serialization and deserialization.  The wire protocol makes up the lowest level in the architecture of a bitcoin node, so the protocol package was designed to function as a standalone library with minimal dependencies for maximal reuse.