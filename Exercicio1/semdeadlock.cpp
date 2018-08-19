#include <iostream>
#include <thread>
#include <mutex>
#include <fstream>
#include <vector>
#include <string>
#include <sstream>

// Variáveis que serão manipuladas pelas threads
int good_quantity = 0;
int index = 0;
int repair_quantity = 0;

// Mutexes
std::mutex mutex_files,mutex_bom;
std::mutex mutex_array[100];


void contador(std::vector<int>& flag, std::vector<int>& repair_quantity ,int num_arch){
    int arrow_type,quantity;
    // Entrando em região crítica
    mutex_files.lock();
    std::ifstream input;
    while(index < num_arch){
        // Achar um arquivo que nenhuma thread tenha usado
        // e aí abre o arquivo
        while(!input.is_open()){
            std::cout << "arquivo nao aberto, thread : " << std::this_thread::get_id() << std::endl;
            if(!flag[index]){
                flag[index] = 1;
                input.open(std::to_string(index) + ".in");
            }
            index++;
        }
        // Saíndo de região crítica
        std::cout << "arquivo aberto, index :  " << index <<  ", thread : " << std::this_thread::get_id() << " destravando mutex dos arquivos" << std::endl;
        mutex_files.unlock();    
        // Calculando custos a partir do arquivo
        // Duas regiões críticas
        while(input >> arrow_type >> quantity){
            if(quantity >= 0){
                mutex_bom.lock();
                good_quantity += quantity;
                mutex_bom.unlock();
            }
            else{
                mutex_array[arrow_type].lock();
                repair_quantity[arrow_type] -= quantity;
                mutex_array[arrow_type].unlock();
            }
        }
        // travando mutex e fechando arquivo para abrir outro
        mutex_files.lock();
        if(input.is_open()){
            input.close();
        }
    }
    mutex_files.unlock();
}

int main(){
    // Variáveis de quantidade
    int num_threads, arrows, i, num_arch;

    // Lendo arquivo inicial das constantes
    std::string entrada;
    std::ifstream input("entrada.in");
    input >> num_arch >> num_threads >> arrows;
    // Criando estruturas para cálculo dos preços
    std::vector<int> price(arrows, 0);
    std::vector<int> repair_quantity(arrows, 0);
    std::vector<std::thread> threads;
    std::vector<int> flag(num_arch , 0);

    // Lendo arquivo de preços
    for(int i = 0; i < arrows ; i++){
        input >> price[i];
    }

    // Criando threads
    for(int i = 0; i < num_threads; i++){
        threads.push_back(std::thread(contador,std::ref(flag),std::ref(repair_quantity),num_arch));
    }
    // Esperando que terminem
    for(auto& th : threads){
        th.join();
    }

    std::cout << good_quantity << " flechas em bom estado\n\n";
    for(int i = 0 ; i < arrows ; i++){
        std::cout << "Custo de consertar as arrows de tipo " << i+1 << ": R$ " << price[i]*repair_quantity[i]<< std::endl;
    }
    return 0;
}
