#include <iostream>
#include <thread>
#include <mutex>
#include <fstream>
#include <vector>
#include <string>
#include <sstream>

// Variáveis que serão manipuladas pelas threads
int custo_bom = 0;
int indice = 0;
int custo_ruim = 0;

// Mutexes
std::mutex xflags,xbom;
std::mutex mutex_array[100];


void contador(std::vector<int>& flag, std::vector<int>& custo_ruim ,int num_arch){
    int f,q;
    // Entrando em região crítica
    xflags.lock();
    std::ifstream input;
    while(indice < num_arch){
        // Achar um arquivo que nenhuma thread tenha usado
        // e aí abre o arquivo
        while(!input.is_open()){
            std::cout << "Não aberto, thread :" << std::this_thread::get_id() << std::endl;
            if(!flag[indice]){
                flag[indice] = 1;
                input.open(std::to_string(indice) + ".in");
            }
            indice++;
        }
        // Saíndo de região crítica
        std::cout << "aberto, indice :  " << indice <<  ", thread :" << std::this_thread::get_id() << " tentando travar" << std::endl;
        xflags.unlock();    
        // Calculando custos a partir do arquivo
        // Duas regiões críticas
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
    // Variáveis de quantidade
    int num_threads, flechas, i, f, q, rc, num_arch;
    // Lendo arquivo inicial das constantes
    std::string entrada;
    std::ifstream input("entrada.in");
    input >> num_arch >> num_threads >> flechas;
    // Criando estruturas para cálculo dos preços
    std::vector<int> preco(flechas, 0);
    std::vector<int> custo_ruim(flechas, 0);
    std::vector<std::thread> threads;
    std::vector<int> flag(num_arch , 0);
    // Lendo arquivo de preços
    for(int i = 0; i < flechas ; i++){
        input >> preco[i];
    }
    // Criando threads
    for(int i = 0; i < num_threads; i++){
        threads.push_back(std::thread(contador,std::ref(flag),std::ref(custo_ruim),num_arch));
    }
    // Esperando que terminem
    for(auto& th : threads){
        th.join();
    }
    std::cout << custo_bom << " Flechas em bom estado\n\n";
    for(int i = 0 ; i < flechas ; i++){
        std::cout << "Custo de consertar as flechas de tipo " << i+1 << ": R$ " << preco[i]*custo_ruim[i]<< std::endl;
    }
    return 0;
}
