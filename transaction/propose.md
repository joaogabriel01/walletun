what is the idea behind the microsservice?
simple, it has 2 purposes:
    * receive data from api gateway (https) and publish it to a queue to handle potential request bottleneck
    * Subscribe to the same queue to:
        * Save the transaction status in ScyllaDB.
        * Send the request to the Wallet Service using gRPC.
        * Update the status to "Finished" after the process is complete.
