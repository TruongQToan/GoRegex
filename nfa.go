package main

type Link struct {
	Value     byte // '' for epsilon link
	NextState *State
}

type State struct {
	Id      int
	Out1    *Link
	Out2    *Link
	IsFinal bool
}

type StateFrag struct {
	StartState *State
	FinalState *State
}

func compile(regex string) *StateFrag {
	/* compile regex to nfa
	   regex in the postfix form
	   . concatenate
	   | alternative
	   * zero or more
	   ex: (ac|b)* has the form ac.b|*
	*/
	postfix := in2post(regex)
	stateFragStack := make([]*StateFrag, 0)
	id := 0
	for _, c := range []byte(postfix) {
		n := len(stateFragStack)
		if c == '.' {
			stateFrag1 := stateFragStack[n-1]
			stateFrag2 := stateFragStack[n-2]
			stateFragStack = stateFragStack[:n-2]

			e := &Link{
				Value:     0,
				NextState: stateFrag1.StartState,
			}

			stateFrag2.FinalState.Out1 = e
			stateFrag2.FinalState.IsFinal = false
			sf := &StateFrag{
				StartState: stateFrag2.StartState,
				FinalState: stateFrag1.FinalState,
			}

			stateFragStack = append(stateFragStack, sf)

		} else if c == '*' {
			n := len(stateFragStack)
			lastStateFrag := stateFragStack[n-1]
			stateFragStack = stateFragStack[:n-1]
			f := &State{
				Id:      id,
				Out1:    nil,
				Out2:    nil,
				IsFinal: true,
			}

			id++

			e1 := &Link{
				Value:     0,
				NextState: lastStateFrag.StartState,
			}

			e2 := &Link{
				Value:     0,
				NextState: f,
			}

			s := &State{
				Id:      id,
				Out1:    e1,
				Out2:    e2,
				IsFinal: false,
			}

			id++

			e3 := &Link{
				Value:     0,
				NextState: f,
			}

			lastStateFrag.FinalState.Out1 = e3

			e4 := &Link{
				Value:     0,
				NextState: lastStateFrag.StartState,
			}

			lastStateFrag.FinalState.Out2 = e4

			sf := &StateFrag{
				StartState: s,
				FinalState: f,
			}
			stateFragStack = append(stateFragStack, sf)

		} else if c == '|' {
			stateFrag1 := stateFragStack[n-1]
			stateFrag2 := stateFragStack[n-2]
			stateFragStack = stateFragStack[:n-2]

			e1 := &Link{
				Value:     0,
				NextState: stateFrag1.StartState,
			}

			e2 := &Link{
				Value:     0,
				NextState: stateFrag2.StartState,
			}

			s := &State{
				Id:      id,
				Out1:    e1,
				Out2:    e2,
				IsFinal: false,
			}

			id++

			f := &State{
				Id:      id,
				Out1:    nil,
				Out2:    nil,
				IsFinal: true,
			}

			id++

			e3 := &Link{
				Value:     0,
				NextState: f,
			}

			stateFrag1.FinalState.Out1 = e3
			stateFrag2.FinalState.Out1 = e3

			sf := &StateFrag{
				StartState: s,
				FinalState: f,
			}

			stateFragStack = append(stateFragStack, sf)

		} else {
			f := &State{
				Id:      id,
				Out1:    nil,
				Out2:    nil,
				IsFinal: true,
			}

			id++

			link := &Link{
				Value:     byte(c),
				NextState: f,
			}

			s := &State{
				Id:      id,
				Out1:    link,
				Out2:    nil,
				IsFinal: false,
			}

			id++

			sf := &StateFrag{
				StartState: s,
				FinalState: f,
			}

			stateFragStack = append(stateFragStack, sf)

		}
	}

	return stateFragStack[0]

}

func move(states []*State, c byte) []*State {
	result := make([]*State, 0)
	stack := make([]*State, 0)
	added := make(map[int]struct{})

	for _, state := range states {
		stack = append(stack, state)
	}

	for len(stack) > 0 {
		t := stack[0]
		stack = stack[1:len(stack)]

		if t.Out1 != nil && t.Out1.Value == c {
			u := t.Out1.NextState
			if _, ok := added[u.Id]; !ok {
				added[u.Id] = struct{}{}
				result = append(result, u)
			}
		}

		if t.Out2 != nil && t.Out1.Value == c {
			u := t.Out2.NextState
			if _, ok := added[u.Id]; !ok {
				added[u.Id] = struct{}{}
				result = append(result, u)
			}
		}
	}

	return result
}

func epsilonClosure(states []*State) []*State {
	result := make([]*State, 0)
	stack := make([]*State, 0)
	added := make(map[int]struct{})

	for _, state := range states {
		result = append(result, state)
		stack = append(stack, state)
		added[state.Id] = struct{}{}
	}

	for len(stack) > 0 {
		t := stack[0]
		stack = stack[1:len(stack)]

		if t.Out1 != nil && t.Out1.Value == 0 {
			u := t.Out1.NextState
			if _, ok := added[u.Id]; !ok {
				added[u.Id] = struct{}{}
				stack = append(stack, u)
				result = append(result, u)
			}
		}

		if t.Out2 != nil && t.Out2.Value == 0 {
			u := t.Out2.NextState
			if _, ok := added[u.Id]; !ok {
				added[u.Id] = struct{}{}
				stack = append(stack, u)
				result = append(result, u)
			}
		}
	}

	return result
}

func simulate(input string, nfa *State) bool {
	visit := make([]*State, 0)

	s0Closure := epsilonClosure([]*State{nfa})
	for _, state := range s0Closure {
		visit = append(visit, state)
	}

	for _, c := range input {
		nextStates := move(visit, byte(c))
		closure := epsilonClosure(nextStates)
		visit = visit[:0]
		for _, state := range closure {
			visit = append(visit, state)
		}
	}

	for _, state := range visit {
		if state.IsFinal {
			return true
		}
	}
	return false
}
