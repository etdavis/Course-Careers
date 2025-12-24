/* Question 5
ABC corporation has many employees, some of which are managers, others which are
employees. Each employee has a manager, and some managers have a manager who
oversees them. Each employee knows who manages them but unfortunately the managers
have lost their org chart making it difficult for them to remember which employees are below
them in the hierarchy.
For this question you’ll be given an array of javascript objects representing employee
information. Each employee has a unique “id” property, and every employee that has a manager
has a “manager” property that stores the “id” of their manager. Your job is to return a new
javascript object that contains the ids of managers as properties, where each property stores an
array of the unique ids of all employees beneath that manager in the hierarchy. You may
assume any manager referenced will always exist in the array. See the sample inputs for further
explanation. */

/*
function createHierarchy(employees) {
    let hierarchy = {};
    for(let employee of employees) {
        if("manager" in employee) {
            if(!(employee.manager in hierarchy)) {
                hierarchy[employee.manager] = [];
            }
            hierarchy[employee.manager].push(employee.id);
            
        }
    }
}
*/
//Got this far and couldn't figure out how to get employees that are indirect reports so I consulted the solution to learn from

function buildHierarchy(employees) {
    let hierarchy = {};
    // Initialize hierarchy
    for (let employee of employees) {
        if (!hierarchy[employee.id]) {
            hierarchy[employee.id] = [];
        }
        if (employee.manager) {
            if (!hierarchy[employee.manager]) {
                hierarchy[employee.manager] = [employee.id];
            } else {
                hierarchy[employee.manager].push(employee.id);
            }
        }
    }
    // This part of the solution I basically had already but slightly different so my start wouldn't work with the rest of the solution

    // Function to get all indrect reports of a manager
    function getReports(manager) {
        let reports = [];
        let queue = [...hierarchy[manager]]; // Stores the inital array of direct reports
        while (queue.length > 0) {
            let current = queue.shift(); // Pops the next employee from the queue
            reports.push(current); // Add the current employee to the array being updated and returned
            queue.push(...hierarchy[current]); // Add the direct reports of the current employee to the queue
        } //Ends when the queue dries up when the last employees have no direct reports
        return reports; // Returns array now with direct and indirect reports
    }

    // Build the final hierarchy with indirect reports
    let finalHierarchy = {};
    for (let manager in hierarchy) {
        if (hierarchy[manager].length > 0) { // If the manager has direct reports
            finalHierarchy[manager] = getReports(manager); // Get all reports, direct and indirect, for that manager
        }
    }
    return finalHierarchy;
}

const employeeData = [
    {"id": "1", "manager": "2"},
    {"id": "2", "manager": "3"},
    {"id": "3"},
    {"id": "4", "manager": "3"}
]
console.log(buildHierarchy(employeeData));
// {
// '2': ['1'],
// '3': ['2', '4', '1']
// }