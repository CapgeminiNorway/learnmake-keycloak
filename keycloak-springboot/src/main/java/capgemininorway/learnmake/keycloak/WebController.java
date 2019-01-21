package capgemininorway.learnmake.keycloak;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

import java.security.Principal;

import org.springframework.beans.factory.annotation.Autowired;

@Controller
public class WebController {

  @Autowired
  private CustomerDAO customerDAO;

  @GetMapping(path = "/")
  public String index() {
    return "external";
  }

  @GetMapping(path = "/customers")
  public String customers(Principal principal, Model model) {
    addCustomers();
    Iterable<Customer> customers = customerDAO.findAll();
    model.addAttribute("customers", customers);
    model.addAttribute("username", principal.getName());
    return "customers";
  }

  // add customers for demonstration
  public void addCustomers() {

    Customer customer;
    //List<Customer> customers = new LinkedList<>();

    customer = new Customer();
    customer.setAddress("Address etc etc -01");
    customer.setName("customer -01");
    customer.setServiceRendered("blabla services -01");
    customerDAO.save(customer);

    customer = new Customer();
    customer.setAddress("Address etc etc -02");
    customer.setName("customer -02");
    customer.setServiceRendered("blabla services -02");
    customerDAO.save(customer);


  }
}
