package base

import (
    "encoding/binary"
    "bytes"
)

type PduHeader struct {
    length     int32
    version    int16
    flag       int16
    serviceId int16
    commandId int16
    seqNum    int16
    reversed   int16
}

func writeUInt32(b []byte,index int,value uint32) {
    _ = b[3 + index]
    b[0 + index] = byte(value >> 24)
    b[1 + index] = byte(value >> 16)
    b[2 + index] = byte(value >> 8)
    b[3 + index] = byte(value)
}


func writeInt32(b []byte,index int,value int32) {
    _ = b[3 + index]
    b[0 + index] = byte(value >> 24)
    b[1 + index] = byte(value >> 16)
    b[2 + index] = byte(value >> 8)
    b[3 + index] = byte(value)
}

func writeInt16(b []byte,index int,value int16) {
    _ = b[1 + index]
    b[index] = byte(value >> 8)
    b[1 + index] = byte(value)
}



type Pdu struct {
    header     *PduHeader
    bufferData []byte
}

func NewPdu() *Pdu {
   pdu := new(Pdu)
   pdu.header = new(PduHeader)
   pdu.bufferData = make([]byte,16)
   return pdu
}

func (this *Pdu)Write(data []byte) {
    this.bufferData = append(this.bufferData,data...)
    this.writeHeader()
}

func (this *Pdu)writeHeader() {
    this.header.length = int32(len(this.bufferData))
    writeInt32(this.bufferData, 0 ,this.header.length);
    writeInt16(this.bufferData, 4 ,this.header.version);
    writeInt16(this.bufferData, 6 ,this.header.flag);
    writeInt16(this.bufferData, 8 ,this.header.serviceId);
    writeInt16(this.bufferData, 10 ,this.header.commandId);
    writeInt16(this.bufferData, 12 ,this.header.seqNum);
    writeInt16(this.bufferData, 14,this.header.reversed);
}

func (this *Pdu)GetBufferData() []byte{
    return this.bufferData
}

func (this *Pdu)GetBodyData()[]byte {
    return this.bufferData[16:]
}

func (this *Pdu)SetServiceId(serviceId int16) {
    this.header.serviceId = serviceId
}

func (this *Pdu)GetServiceId() int16 {
    return this.header.serviceId
}

func (this *Pdu)SetCommandId(commandId int16) {
    this.header.commandId = commandId
}

func (this *Pdu)GetCommandId() int16 {
    return this.header.commandId
}

func (this *Pdu)SetSeqNum(seqNum int16) {
    this.header.seqNum = seqNum
}

func (this *Pdu)GetSeqNum() int16{
    return this.header.seqNum
}

func ReadPdu(buffer []byte) *Pdu {
    bufferSize := int32(len(buffer))
    if(bufferSize < 16) {
        return nil
    }
    pdu := new(Pdu)
    pdu.header = new(PduHeader)
    reader := bytes.NewReader(buffer)
    binary.Read(reader, binary.BigEndian, &pdu.header.length)
    binary.Read(reader, binary.BigEndian, &pdu.header.version)
    binary.Read(reader, binary.BigEndian, &pdu.header.flag)
    binary.Read(reader, binary.BigEndian, &pdu.header.serviceId)
    binary.Read(reader, binary.BigEndian, &pdu.header.commandId)
    binary.Read(reader, binary.BigEndian, &pdu.header.seqNum)
    binary.Read(reader, binary.BigEndian, &pdu.header.reversed)
    if(pdu.header.length > bufferSize) {
        return nil;
    }
    pdu.bufferData = make([]byte, 0 ,16)
    pdu.bufferData = append(pdu.bufferData, buffer...)
    return pdu
}


type AttachData struct {
    Type    uint32
    Handle  uint32
    ServiceType uint32
    Data    []byte
}


func NewAttachData(typeV uint32, handle uint32, serviceType uint32) *AttachData {
    attachData := &AttachData{Type: typeV, Handle: handle, ServiceType: serviceType}
 //   attachData.Serialization()
    return attachData
}


func NewAttachDataForData(data []byte) *AttachData{
   attachData := &AttachData{} 
   reader := bytes.NewReader(data);
   binary.Read(reader, binary.BigEndian, &attachData.Type)
   binary.Read(reader, binary.BigEndian, &attachData.Handle)
   binary.Read(reader, binary.BigEndian, &attachData.ServiceType)
   return attachData
}

func (this *AttachData)Serialization() {
    this.Data = make([]byte,12)
    writeUInt32(this.Data, 0, this.Type)
    writeUInt32(this.Data, 4, this.Handle)
    writeUInt32(this.Data, 8, this.ServiceType)
}

func (this *AttachData)GetBufferData() []byte{
    this.Serialization()
    return this.Data
}


func (this *AttachData)GetHandle() uint32 {
    return this.Handle
}



