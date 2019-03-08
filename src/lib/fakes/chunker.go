// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"lib/policy_client"
	"sync"
)

type Chunker struct {
	ChunkStub        func(allPolicies []policy_client.PolicyV0) [][]policy_client.PolicyV0
	chunkMutex       sync.RWMutex
	chunkArgsForCall []struct {
		allPolicies []policy_client.PolicyV0
	}
	chunkReturns struct {
		result1 [][]policy_client.PolicyV0
	}
	chunkReturnsOnCall map[int]struct {
		result1 [][]policy_client.PolicyV0
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Chunker) Chunk(allPolicies []policy_client.PolicyV0) [][]policy_client.PolicyV0 {
	var allPoliciesCopy []policy_client.PolicyV0
	if allPolicies != nil {
		allPoliciesCopy = make([]policy_client.PolicyV0, len(allPolicies))
		copy(allPoliciesCopy, allPolicies)
	}
	fake.chunkMutex.Lock()
	ret, specificReturn := fake.chunkReturnsOnCall[len(fake.chunkArgsForCall)]
	fake.chunkArgsForCall = append(fake.chunkArgsForCall, struct {
		allPolicies []policy_client.PolicyV0
	}{allPoliciesCopy})
	fake.recordInvocation("Chunk", []interface{}{allPoliciesCopy})
	fake.chunkMutex.Unlock()
	if fake.ChunkStub != nil {
		return fake.ChunkStub(allPolicies)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.chunkReturns.result1
}

func (fake *Chunker) ChunkCallCount() int {
	fake.chunkMutex.RLock()
	defer fake.chunkMutex.RUnlock()
	return len(fake.chunkArgsForCall)
}

func (fake *Chunker) ChunkArgsForCall(i int) []policy_client.PolicyV0 {
	fake.chunkMutex.RLock()
	defer fake.chunkMutex.RUnlock()
	return fake.chunkArgsForCall[i].allPolicies
}

func (fake *Chunker) ChunkReturns(result1 [][]policy_client.PolicyV0) {
	fake.ChunkStub = nil
	fake.chunkReturns = struct {
		result1 [][]policy_client.PolicyV0
	}{result1}
}

func (fake *Chunker) ChunkReturnsOnCall(i int, result1 [][]policy_client.PolicyV0) {
	fake.ChunkStub = nil
	if fake.chunkReturnsOnCall == nil {
		fake.chunkReturnsOnCall = make(map[int]struct {
			result1 [][]policy_client.PolicyV0
		})
	}
	fake.chunkReturnsOnCall[i] = struct {
		result1 [][]policy_client.PolicyV0
	}{result1}
}

func (fake *Chunker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.chunkMutex.RLock()
	defer fake.chunkMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Chunker) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ policy_client.Chunker = new(Chunker)
