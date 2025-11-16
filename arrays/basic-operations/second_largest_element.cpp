// { Driver Code Starts
/* Driver program to test above function */

#include<bits/stdc++.h>
using namespace std;

 // } Driver Code Ends
class Solution
{
   public:
    int print2largest(int arr[], int arr_size)
    {
    	int first_max = arr[0],second_max = INT_MIN;
    	for(int i=0;i<arr_size;i++)
    	{
    	    if(arr[i] > first_max)
    	    {
    	        second_max = first_max;
    	        first_max = arr[i];     
    	    }
    	    
    	    else if(arr[i] > second_max && arr[i] < second_max)
    	        second_max = arr[i];
    	}
    	
    	return second_max;
    }

};

// { Driver Code Starts.
int main()
{
	int t;
	cin>>t;
	while(t--)
	{
	    int n;
	    cin>>n;
	    int arr[n];
	    for(int i=0;i<n;i++)
	      cin>>arr[i];
	    Solution ob;  
	    int ans=ob.print2largest(arr, n);
	    cout<<ans;
	    cout<<"\n";
	}
	return 0;
}
  // } Driver Code Ends