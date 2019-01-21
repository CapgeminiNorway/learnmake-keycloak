package capgemininorway.learnmake.keycloak;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

@ToString @Getter @Setter
@Entity
public class Customer {

  @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
  private long id;

  private String name;
  private String serviceRendered;
  private String address;

}
