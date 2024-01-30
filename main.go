package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"bytes"
	"github.com/btcsuite/btcd/wire"
)

func main() {
	// Check if the user provided a raw transaction hex input
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <raw_transaction_hex>")
		return
	}

	// Parse the raw transaction hex
	rawTxHex := os.Args[1]
	rawTxBytes, err := hex.DecodeString(rawTxHex)
	if err != nil {
		fmt.Println("Error decoding raw transaction hex:", err)
		return
	}

	// Deserialize the raw transaction bytes
	var tx wire.MsgTx
	err = tx.Deserialize(bytes.NewReader(rawTxBytes))
	if err != nil {
		fmt.Println("Error deserializing raw transaction:", err)
		return
	}

	// Print transaction details
	fmt.Println("Transaction Version:", tx.Version)
	fmt.Println("Transaction Locktime:", tx.LockTime)
	fmt.Println("Number of Inputs:", len(tx.TxIn))
	fmt.Println("Number of Outputs:", len(tx.TxOut))

	// Print input details
	fmt.Println("\nInputs:")
	for i, input := range tx.TxIn {
		fmt.Printf("Input %d:\n", i)
		fmt.Printf("  Previous Tx Hash: %s\n", input.PreviousOutPoint.Hash)
		fmt.Printf("  Previous Tx Index: %d\n", input.PreviousOutPoint.Index)
		fmt.Printf("  Script Length: %d\n", len(input.SignatureScript))
		fmt.Println("  Script:", hex.EncodeToString(input.SignatureScript))
	}

	// Print output details
	fmt.Println("\nOutputs:")
	for i, output := range tx.TxOut {
		fmt.Printf("Output %d:\n", i)
		fmt.Printf("  Value: %d Satoshis\n", output.Value)
		fmt.Printf("  Script Length: %d\n", len(output.PkScript))
		fmt.Println("  Script:", hex.EncodeToString(output.PkScript))
	}
}

