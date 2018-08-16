#include <iostream>
#include <thread>
#include <mutex>
#include <fstream>
#include <vector>
#include <string>
#include <sstream>

int custo_bom = 0;
int indice = 0;
int custo_ruim = 0;

std::mutex xflags,xbom;
std::mutex mutex_array[100];


void contador(std::vector<int>& flag, std::vector<int>& custo_ruim ,int num_arch){
    int f,q;
    xflags.lock();
    std::ifstream input;
    while(indice < num_arch){
        while(!input.is_open()){
            std::cout << "Não aberto, thread :" << std::this_thread::get_id() << std::endl;
            if(!flag[indice]){
                flag[indice] = 1;
                input.open(std::to_string(indice) + ".in");
            }
            indice++;
        }
        std::cout << "aberto, indice :  " << indice <<  ", thread :" << std::this_thread::get_id() << " tentando travar" << std::endl;
        xflags.unlock();    
        std::cout << "travado, calculando" << std::this_thread::get_id() << std::endl;
        while(input >> f >> q){
            if(q == 0){
                custo_bom += q;
            }
            else if(q > 0){
                xbom.lock();
                custo_bom += q;
                xbom.unlock();
            }
            else{
                mutex_array[f].lock();
                custo_ruim[f] -= q;
                mutex_array[f].unlock();
            }
        }
        xflags.lock();
        if(input.is_open()){
            input.close();
        }
    }
    xflags.unlock();
}

int main(){
    int num_threads, flechas, i, f, q, rc, num_arch;
    std::string entrada;
    std::ifstream input("entrada.in");
    input >> num_arch >> num_threads >> flechas;
    std::vector<int> preco(flechas, 0);
    std::vector<int> custo_ruim(flechas, 0);
    std::vector<std::thread> threads;
    std::vector<int> flag(num_arch , 0);
    for(int i = 0; i < flechas ; i++){
        input >> preco[i];
    }
    for(int i = 0; i < num_threads; i++){
        threads.push_back(std::thread(contador,std::ref(flag),std::ref(custo_ruim),num_arch));
    }
    for(auto& th : threads){
        th.join();
    }
    std::cout << custo_bom << " Flechas em bom estado\n\n";
    for(int i = 0 ; i < flechas ; i++){
        std::cout << "Custo de consertar as flechas de tipo " << i+1 << ": R$ " << preco[i]*custo_ruim[i]<< std::endl;
    }
    return 0;
}
