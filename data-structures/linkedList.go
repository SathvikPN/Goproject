package datastructures

type Node struct {
    data int
    next *Node
}

type MyLinkedList struct {
    head *Node
    len int
}

func Constructor() MyLinkedList {
    return MyLinkedList{}
}


func (this *MyLinkedList) Get(index int) int {
    if !(index < this.len) {
        return -1
    }

    current := this.head
    for i:=1; i<=index; i++ {
        current = current.next
    }

    return current.data
}


func (this *MyLinkedList) AddAtHead(val int)  {
    newNode := &Node{val, this.head}
    this.head = newNode
    this.len++
}


func (this *MyLinkedList) AddAtTail(val int)  {
    if this.len == 0 {
        this.AddAtHead(val)
        return
    }

    previousNode := this.head
    for previousNode.next != nil {
        previousNode = previousNode.next
    }
    previousNode.next = &Node{val, nil}
    this.len++
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index > this.len {
        return 
    }
    if index == this.len {
        this.AddAtTail(val)
        return
    }
    if index == 0 {
        this.AddAtHead(val)
        return
    }
    previousNode := this.head
    for i:=1; i<index; i++ {
        previousNode = previousNode.next
    }
    newNode := &Node{val, previousNode.next}
    previousNode.next = newNode
    this.len++
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if !(index < this.len) {
        return
    }

    if index == 0 {
        this.head = this.head.next
        this.len--
        return
    }

    previousNode := this.head
    for i:=1; i<index; i++ {
        previousNode = previousNode.next
    }
    outNode := previousNode.next
    previousNode.next = outNode.next
    this.len--
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
