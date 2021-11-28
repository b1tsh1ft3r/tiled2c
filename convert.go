package main

import "log"
import "bufio"
import "encoding/csv"
import "os"
import "fmt"
import "strconv"

func main() {
    fmt.Println("**************************************");
    fmt.Println("TILED MAP EDITOR CSV EXPORT TO C ARRAY");
    fmt.Println("**************************************\n");

    var FILENAME string
    var TILE_SIZE, HEIGHT_TILES, WIDTH_TILES, MAP_WIDTH, MAP_HEIGHT int

    fmt.Println("Please enter the filename of the exported CSV Tilemap")
    fmt.Scan(&FILENAME);

    INPUT_FILE, ERR  := os.Open(FILENAME);
    if ERR != nil { log.Fatal(ERR); }

    fmt.Println("TILE SIZE? (Ex: 8,16,24,32)");
    fmt.Scan(&TILE_SIZE);

    fmt.Println("MAP HEIGHT IN TILES?");
    fmt.Scan(&HEIGHT_TILES);

    fmt.Println("MAP WIDTH IN TILES?");
    fmt.Scan(&WIDTH_TILES);

    MAP_WIDTH=(WIDTH_TILES*TILE_SIZE);
    MAP_HEIGHT=(HEIGHT_TILES*TILE_SIZE);

    fmt.Println("\n\nCONVERTING TILE MAP...\n");

    CSV_MAP          := csv.NewReader(bufio.NewReader(INPUT_FILE));
    TILEMAP_DATA, _  := CSV_MAP.ReadAll();  // READ ALL DATA FROM CSV VERSION OF TILEMAP (EXPORTED FROM TILED MAP EDITOR)
    OUTPUT_FILE, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644);
    if err != nil { log.Fatal(err); }
    
    OUTPUT_FILE.WriteString("const u8 level1["+strconv.Itoa(MAP_HEIGHT)+"]["+strconv.Itoa(MAP_WIDTH)+"] = {\n");

    for i:=0;i<=HEIGHT_TILES-1;i++ {

        for j:=0;j<=WIDTH_TILES-1;j++ {

            tile_value := ((TILEMAP_DATA[i])[j]);
            if (j == 0) { OUTPUT_FILE.WriteString("     { "+tile_value+", "); continue; }
            if (j == WIDTH_TILES-1 && i >= HEIGHT_TILES-1) { OUTPUT_FILE.WriteString(tile_value+" }\n};\n"); continue; }            
            if (j == WIDTH_TILES-1) { OUTPUT_FILE.WriteString(tile_value+" },\n"); continue; }            
            OUTPUT_FILE.WriteString(tile_value+", ");

        }
    }

    if err := OUTPUT_FILE.Close(); err != nil { print(err); }
    fmt.Println("TILEMAP CONVERTED!");

}