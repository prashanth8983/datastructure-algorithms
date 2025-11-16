#include <iostream>
using namespace std;

int main() {
	string A;
	cin>>A;
	
	for(int i=0;i<A.size();i++)
	{
	    for(int j=i+1;j<A.size()-i;i++)
	        cout<<A[j];
	}
	
	return 0;
}