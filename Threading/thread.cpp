 #include <chrono>   //for sleeping the thread and making it delay ; 
#include <iostream>
#include <thread>
#include <mutex>

// declearing mutex here; 
std:: mutex mtx; 
void sayhello(){
  std::cout<<"hello from new thread!"<<std::endl; 
}
// real shit that i want to do with my thread;  
 void task2(){
   std::cout<<"this is done by my second thread\n";
   std::this_thread::sleep_for(std::chrono::seconds(2));
   std::cout<<"task 1 finished";
 }
 void  task3(){
   std::cout<<"hello from third thread";
   std::this_thread::sleep_for(std::chrono::seconds(3));
   std::cout<<"task 3";
 }
 
 // generating rat-race condition ;
 // fixing this condition ; 
int task4(int n){
  int total_counter=0; 
  std::lock_guard<std::mutex>lock(mtx);
  for(int i =0; i<=n ; i++){
    total_counter+=1; 
  }
  
  // mutux is already going to unlock when locks goes of scope 
  return total_counter; 
} 


 
 
int main(){
  // creating a new thread; 
  std::thread t(sayhello);
  std::thread t2(task2);
  std::thread t3(task3);
  // main thread waits for completing the t threading 
  // creating rat race multithreading condition
  // means when 2 thread tries to access the same data at same time ; this
  // can be prevent by locking which is called mutex; 
  
  // t4 is going to read and all thread are independent they dont have to be dependent on each other in complied language; 
  std::thread t4(task4);
  std::thread t5(task4);

  t.join();
  t2.join();
  t3.join();
  t4.join();
  t5.join();
  std::cout<<"hello from the main thread!"<<std::endl; 
  return 0; 
}