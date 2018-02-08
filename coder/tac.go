package coder

import "fmt"

type TACprogram struct {
    Vars   []string
    Funcs  []TACfunction
}

type TACfunction struct {
    Vars map[string]TACaddr
    Args []string

    Body []TACstatement
}

type TACstatement struct {
    Label    string
    Operator string
    Target   TACaddr
    Arg1     TACaddr
    Arg2     TACaddr
}

const (
    TAC_VARIABLE   = iota
    TAC_TEMP       = iota
    TAC_INTLITERAL = iota
    TAC_ARGUMENT   = iota
    TAC_NONE       = iota
    TAC_RETRUN     = iota
)

type TACaddr struct {
    TACtype int
    Index   int
}

func (tfunc TACfunction) String() string {
    str := ""
    for _, stmt := range tfunc.Body {
        str += stmt.String() + "\n"
    }
    return str
}

func (tstmt TACstatement) String() string {
    return fmt.Sprintf("%10s %6s := %s (%s, %s)",
        tstmt.Label, tstmt.Operator,
        tstmt.Target.String(), tstmt.Arg1.String(),
        tstmt.Arg2.String(),
    )
}

func (taddr TACaddr) String() string {
    if taddr.TACtype == TAC_VARIABLE {
        return fmt.Sprintf("v%d", taddr.Index)
    } else if taddr.TACtype == TAC_TEMP {
        return fmt.Sprintf("t%d", taddr.Index)
    } else if taddr.TACtype == TAC_RETURN {
        return fmt.Sprintf("r%d", taddr.Index)
    } else if taddr.TACtype == TAC_INTLITERAL {
        return fmt.Sprintf("%d", taddr.Index)
    } else if taddr.TACtype == TAC_NONE {
        return ""
    }
    return "ERROR!"
}
