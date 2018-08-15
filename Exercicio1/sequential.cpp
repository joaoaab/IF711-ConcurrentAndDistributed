#include <iostream>
#include <thread>
#include <mutex>
#include <fstream>
#include <vector>

int custom_bom = 0;
int indice = 0;

std::mutex xflags,xbom;

template <typename T>
  std::string NumberToString ( T Number )
  {
     ostringstream ss;
     ss << Number;
     return ss.str();
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
    int aux;
    for(int i = 0; i < flechas ; i++){
        input >> preco[i];
    }

    return 0;
}



void contador(std::vector<int>& flag, int num_arch){
    int f,q;
    xflags.lock();
    std::ifstream input;
    if(input){
        std::cout << "wow" << std::endl;
        return;
    }
    while(indice < num_arch){
        xbom.lock();
        if(!flag[indice]){
            flag[indice] = 1;
            input.open(NumberToString(indice) + ".in");
            break;
        }
    }



}
