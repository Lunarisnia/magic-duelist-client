package magicp

// What the client need to transfer
// 1. Their position (SNAPSHOT OF THE WORLD)
// 2. Each bullets starting position they fire (SNAPSHOT OF THE WORLD)
// 3. Their intent (moving, shooting, etc.)

// What the server need to inform
// 1. Each of the players position
// 2. Each bullets position on the field
// 3. If the game is over (eg: player got hit)

// Client Request Protocol:
// {
//   "position": {
//     "x": 100,
//     "y": 0,
//   },
//   "intent": "IDLE"
// }

// Server Snapshot Protocol:
// {
//   "p1_position": {
//     "x": 100,
//     "y": 200
//   },
//   "p2_position": {
//     "x": 300,
//     "y": 0
//   },
//   "bullets": {
//     "owner": "p1",
//     "position": {
//       "x": 0,
//       "y": 100
//     },
//     "direction": {
//       "x": 1,
//       "y": 0
//     },
//     "prev": null,
//     "next": {
//       "owner": "p2",
//       "position": {
//         "x": 0,
//         "y": 1
//       },
//       "direction": {
//         "x": -1,
//         "y": 0
//       }
//     }
//   }
// }
