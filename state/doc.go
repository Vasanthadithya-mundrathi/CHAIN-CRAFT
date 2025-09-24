// Package state provides a structure for chaincraft's ability to access
// state-relevant information from as well as submit transactions/messages
// to the chaincraft network.
//
// This package contains one main interface, `Accessor`, that defines
// the methods available for both accessing and updating state on the
// chaincraft network.
//
// `Accessor` will contain three different implementations:
//  1. Implementation over a gRPC connection with a chaincraft-core node
//     called `CoreAccess`.
//  2. Implementation over a libp2p stream with a state-providing node.
//  3. Implementation over a local running instance of the
//     chaincraft-application (this feature will be implemented in *Full*
//     nodes).
package state
