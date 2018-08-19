#include <iostream>
#include <thread>
#include <mutex>
#include <fstream>
#include <vector>
#include <string>

int good_quantity = 0;
int index = 0;
int repair_quantity = 0;

std::mutex mutex_files,mutex_bom;
std::mutex mutex_array[100];



void contador(std::vector<int>& flag, std::vector<int>& repair_quantity ,int num_arch){
    int arrow_type,quantity;
    mutex_files.lock();
    std::ifstream input;
    while(index < num_arch){
        mutex_bom.lock();
        while(!input.is_open()){
            if(!flag[index]){
                flag[index] = 1;
                input.open(std::to_string(index) + ".in");
            }
            index++;
        }
        mutex_files.unlock();    
        while(input >> arrow_type >> quantity){
            if(q == 0){
                good_quantity += quantity;
            }
            else{
                mutex_array[arrow_type].lock();
                repair_quantity[arrow_type] -= quantity;
                mutex_array[arrow_type].unlock();
            }
        }
        mutex_files.lock();
        mutex_bom.unlock();
        if(input.is_open()){
            input.close();
        }
    }
    mutex_files.unlock();
}



int main(){
    int num_threads, arrows, i, f, q, rc, num_arch;
    std::string entrada;
    std::ifstream input("entrada.in");
    input >> num_arch >> num_threads >> arrows;
    std::vector<int> price(arrows, 0);
    std::vector<int> repair_quantity(arrows, 0);
    std::vector<std::thread> threads;
    std::vector<int> flag(num_arch , 0);
    std::vector<std::mutex> mutex_array(arrows);
    for(int i = 0; i < arrows ; i++){
        input >> price[i];
    }
    for(int i = 0; i < num_threads; i++){
        threads.push_back(std::thread(contador,std::ref(flag), std::ref(repair_quantity),num_arch));
    }
    for(auto& th : threads){
        th.join();
    }
    std::cout << mutex_bom << " arrows em bom estado\n\n";
    for(int i = 0 ; i < arrows ; i++){
        std::cout << "Custo de consertar as arrows de tipo " << i+1 << ": R$ " << price[i]*repair_quantity[i]<< std::endl;
    }
    return 0;
}