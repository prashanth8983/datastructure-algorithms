#include <bits/stdc++.h>

using namespace std;

class Heap {
    private: 
        vector<int> heap;

        int parent(int index){
            return (index-1)/2;
        }

        int leftChild(int index){
            return 2*index+1;
        }

        int rightChild(int index){
            return 2*index+2;
        }

        void heapifyUp(int index){
            while(index > 0 and heap[index]){
                swap(heap[parent(index)], heap[index]);
                index = parent(index);
            }
        }

        void heapifyDown(int index){
            int largest = index;
            int left = leftChild(index);
            int right = rightChild(index);

            if(heap[left] < heap.size() and heap[left] > heap[largest]){
                largest = left;
            }
            if(heap[right] < heap.size() and heap[right] > heap[largest]){
                largest = right;
            }

            if(largest != index){
                swap(heap[index], heap[largest]);
                heapifyDown(largest);
            }
        }
    
    public:
            void insert(int val){
                heap.push_back(val);
                heapifyUp(heap.size()-1);
            }

            int extractMax(){
                if(heap.empty()){
                    throw runtime_error("heap is empty");
                }

                int max_value = heap[0];
                heap[0] = heap.back();
                heap.pop_back();
                if(!heap.empty()){
                    heapifyDown(0);
                }
                return max_value;
            }

            int get_max(){
                if(heap.empty()){
                    throw runtime_error("heap is empty");
                }
                return heap[0];
            }
};

int main(){
    return 0;
}