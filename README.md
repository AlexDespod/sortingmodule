# sortingmodule
This module intended for sorting files which contain a data by type "number\n".
Functional of this module was cteated for sorting a large files which need to be sorted by external sort.
For example : you have a file whose size 100GB , but RAM of your machine limeted up to 1GB .

In this case you should to use this algoritm:

    1) create a folder for chunks 
    2) call function sorter.SortFile(path to file, path to chunk folder, number of records to write a chunk)
    3) open file for output
    4) call function utils.MergeChunks(dir of chunks, outputFile *os.File)
    
    
If file not so large and you have enough RAM for sorting , you can use a async function "sorter.WorkerPoolSort("test.txt","chunks", 5000000, 12)".
In this case you should to calculate some .
  1.(for sorter.WorkerPoolSort)
    1) first argument is a inFile
    2) second argument is a chunk dir
    3) third argument is a number of records to write a chunk
    4) fourth argument mast be ~ int(inFileSize/chunkSize) . It is a number of workers each of which do a chunk
    
  2.
    1) open file for output
    2) call function utils.MergeChunks(dir of chunks, outputFile *os.File)
  
 Function sorter.WorkerPoolSort use a context , thats why , if one of workers occured an error , information about this , will be delivered to other goroutines an pool,
 goroutines will exit.
 
 For merging chunks , i used a Max-heap (<a>https://en.wikipedia.org/wiki/Min-max_heap</a>)
 
 speed test :
              given :
                      file with 60 million records ~ 1GB
              results :
                      async method was faster in 2-2.5 times then siple sort (but async method load your CPU to maximum, and take many resource of RAM)
 
