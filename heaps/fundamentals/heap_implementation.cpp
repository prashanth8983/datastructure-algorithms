#include <bits/stdc++.h>

using namespace std;

class Heap
{
private:
    vector<int> heap;

    int parent(int i)
    {
        return (i - 1) / 2;
    }

    int leftChild(int i)
    {
        return 2 * i + 1;
    }

    int rightChild(int i)
    {
        return 2 * i + 2;
    }

    void heapifyUp(int index)
    {
        while (index > 0 && heap[index] > heap[parent(index)])
        {
            swap(heap[index], heap[parent(index)]);
            index = parent(index);
        }
    }

    void heapifyDown(int index)
    {
        int largest = index;
        int left = leftChild(index);
        int right = rightChild(index);

        if (left < heap.size() && heap[left] > heap[largest])
        {
            largest = left;
        }
        if (right < heap.size() && heap[right] > heap[largest])
        {
            largest = right;
        }

        if (largest != index)
        {
            swap(heap[index], heap[largest]);
            heapifyDown(largest);
        }
    }

public:
    Heap() {}

    void insert(int value){
        heap.push_back(value);
        heapifyUp(heap.size()-1);
    }

    int extractMax(){
        if(heap.empty()){
            throw runtime_error("Heap is empty");
        }

        int maxValue = heap[0];
        heap[0] = heap.back();
        heap.pop_back();
        if(!heap.empty()){
            heapifyDown(0);
        }
        return maxValue;
    }

    int getMax() const {
         if(heap.empty()){
            throw runtime_error("heap is emplty");
         }
         return heap[0];
    }

    bool isEmplty() const {
        return heap.empty();
    }

    int size() const {
        return heap.size();
    }

    
};

int main()
{
    return 0;
}