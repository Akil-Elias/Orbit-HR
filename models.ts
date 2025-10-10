const enum Gender {
    Male = "male",
    Female = "female"
}

class Position {
    title: string;
    department: Department;
    division: Division;
    supervisor: string;
    directReports: Position[];
    
    constructor(
        id: number,
        title: string,
        department: Department,
        division: Division,
        supervisor: string,
        directReports: Position[],
    ){
        this.title = title
    }
}


class Employee {
    constructor(
        id: number,
        //personal info
        firstName: string, 
        lastName: string,
        gender: Gender,
        dob: string,
        age: number, 
        parent: boolean,
        maritalStatus: boolean,
        //job info
        position: string,

    ){

    }
}