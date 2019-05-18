package handlers

import (
	"log"

	"github.com/Easprey/whiteboard-server/models"
	"github.com/Easprey/whiteboard-server/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// operations.<operationID>HandlerFunc expects an anonymous function with the appropriate
// signature to be passed.
// We call <operationID>(db *DataBase) so the handler can consume the database connection.
// <operationID>(db* DataBase) returns the anonymous function <operationID> expects.
// The returned anonymous function has access to the database connection.
// This paradigm is known as a closure.

func GetFingerPaths(db *DataBase) func(operations.FingerPathsGetParams) middleware.Responder {
	return func(params operations.FingerPathsGetParams) middleware.Responder {
		// create response object
		r := operations.NewFingerPathsGetOK()

		// query the database
		q := "SELECT fingerpaths.path_id, path_color, board_color, dash, blur, clear, " +
			"user_id, x, y FROM fingerpaths INNER JOIN fingerpoints WHERE board = ? AND " +
			"fingerpaths.path_id = fingerpoints.path_id"
		rows, err := db.Query(q, params.BoardName)
		if err != nil {
			log.Fatal(err) // TODO: return server error
		}
		defer rows.Close()

		var lastPath models.FingerPath
		for rows.Next() {
			// read row data into local path and point
			var path models.FingerPath
			var point models.FingerPoint
			err := rows.Scan(&path.PathID, &path.PathColor, &path.BoardColor, &path.Dash, &path.Blur,
				&path.Clear, &path.UserID, &point.X, &point.Y)
			if err != nil {
				log.Fatal(err) // TODO: return server error
			}

			if lastPath.PathID == 0 {
				lastPath = path
				lastPath.FingerPoints = append(path.FingerPoints, &point)
			} else if path.PathID == lastPath.PathID {
				// simply append point to lastPath if it is an extension of lastPath
				lastPath.FingerPoints = append(lastPath.FingerPoints, &point)
			} else {
				// r.Payload expects a *FingerPath
				// cannot provide &lastPath, because &lastPath changes with each new found path
				// copy lastPath into new savePath struct and pass &savePath instead
				savePath := lastPath
				r.Payload = append(r.Payload, &savePath)
				// "reset" lastPath and append point
				lastPath = path
				lastPath.FingerPoints = append(path.FingerPoints, &point)
			}
		}

		// handle the final lastPath, which the loop structure above does not append to Payload
		savePath := lastPath
		r.Payload = append(r.Payload, &savePath)

		return r
	}
}

func PostFingerPaths(db *DataBase) func(operations.FingerPathsPostParams) middleware.Responder {
	return func(params operations.FingerPathsPostParams) middleware.Responder {
		// for each fingerpath we receive
		for _, path := range params.Body {
			// begin a new transaction
			tx, err := db.Begin()
			if err != nil {
				log.Fatal(err)
			}
			defer tx.Rollback()

			// prepare SQL statements
			pathStmt, err := tx.Prepare("INSERT INTO fingerpaths (board, path_color, board_color, dash, " +
				"blur, clear, user_id) VALUES (?, ?, ?, ?, ?, ?, ?)")
			if err != nil {
				log.Fatal(err)
			}
			defer pathStmt.Close()
			pointStmt, err := tx.Prepare("INSERT INTO fingerpoints (path_id, x, y) VALUES (?, ?, ?)")
			if err != nil {
				log.Fatal(err)
			}
			defer pointStmt.Close()

			// insert fingerpath into database
			result, err := pathStmt.Exec(params.BoardName, *path.PathColor, *path.BoardColor, *path.Dash, *path.Blur, *path.Clear, path.UserID)
			if err != nil {
				log.Fatal(err)
			}

			// insert associated fingerpoints into database
			for _, point := range path.FingerPoints {
				id, err := result.LastInsertId()
				if err != nil {
					log.Fatal(err)
				}
				_, err = pointStmt.Exec(id, *point.X, *point.Y)
				if err != nil {
					log.Fatal(err)
				}
			}

			// commit database transaction
			err = tx.Commit()
			if err != nil {
				log.Fatal(err)
			}
		}

		return operations.NewFingerPathsPostCreated()
	}
}
