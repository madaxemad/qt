package xml

//#include "xml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QXmlDeclHandler struct {
	ptr unsafe.Pointer
}

type QXmlDeclHandler_ITF interface {
	QXmlDeclHandler_PTR() *QXmlDeclHandler
}

func (p *QXmlDeclHandler) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlDeclHandler) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlDeclHandler(ptr QXmlDeclHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlDeclHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlDeclHandlerFromPointer(ptr unsafe.Pointer) *QXmlDeclHandler {
	var n = new(QXmlDeclHandler)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlDeclHandler_") {
		n.SetObjectNameAbs("QXmlDeclHandler_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlDeclHandler) QXmlDeclHandler_PTR() *QXmlDeclHandler {
	return ptr
}

func (ptr *QXmlDeclHandler) AttributeDecl(eName string, aName string, ty string, valueDefault string, value string) bool {
	defer qt.Recovering("QXmlDeclHandler::attributeDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDeclHandler_AttributeDecl(ptr.Pointer(), C.CString(eName), C.CString(aName), C.CString(ty), C.CString(valueDefault), C.CString(value)) != 0
	}
	return false
}

func (ptr *QXmlDeclHandler) ErrorString() string {
	defer qt.Recovering("QXmlDeclHandler::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDeclHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlDeclHandler) ExternalEntityDecl(name string, publicId string, systemId string) bool {
	defer qt.Recovering("QXmlDeclHandler::externalEntityDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDeclHandler_ExternalEntityDecl(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlDeclHandler) InternalEntityDecl(name string, value string) bool {
	defer qt.Recovering("QXmlDeclHandler::internalEntityDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDeclHandler_InternalEntityDecl(ptr.Pointer(), C.CString(name), C.CString(value)) != 0
	}
	return false
}

func (ptr *QXmlDeclHandler) DestroyQXmlDeclHandler() {
	defer qt.Recovering("QXmlDeclHandler::~QXmlDeclHandler")

	if ptr.Pointer() != nil {
		C.QXmlDeclHandler_DestroyQXmlDeclHandler(ptr.Pointer())
	}
}

func (ptr *QXmlDeclHandler) ObjectNameAbs() string {
	defer qt.Recovering("QXmlDeclHandler::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDeclHandler_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlDeclHandler) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlDeclHandler::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlDeclHandler_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}