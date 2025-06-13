#include <iostream>
#include <unistd.h>
#include <sys/wait.h>
#include <cerrno>
#include <cstring>
#include <cstdlib>

void pipeline(const char *process1, const char *process2) {
  int fd[2];
  pipe(fd);
  int id = fork();
  if (id != 0) {
    // close the read end of file; 
    close(fd[0]);
    // redirect the output to the write file;
    dup2(fd[1], STDOUT_FILENO);
    close(fd[1]);
    // run the first command;
  
    execlp("cat", "cat", "pipeline.cpp", nullptr); // FIXED

    std::cerr << "failed to execute " << process1 << std::endl;
  }
  else {
    // close the write end of pipe;
    close(fd[1]);
    dup2(fd[0], STDIN_FILENO);
    close(fd[0]);
    // executing the process 2; with using exec command;
    execlp("grep", "grep", "int", nullptr); 

    std::cerr << "failed to execute " << process2 << std::endl; 
}
}

int main() {
  pipeline("cat pipeline.cpp", "grep int");
  return 0;
}
