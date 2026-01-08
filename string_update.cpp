#include <iostream>
#include "bits/stdc++.h"

using namespace std;

int main()
{

    string a = "abc";

    unordered_set<string> data = {"cbc", "abd", "xyz"};

    for (int i = 0; i < a.length(); i++)
    {
        
        for (int j = 0; j < 26; j++)
        {
            a[i] = (char)(j+97);

            if(data.find(a) != data.end()){
                cout<<a<<"\n";
            }
        }
    }

    return 0;
}